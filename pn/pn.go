package pn

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/lyraproj/issue/issue"
)

type (
	// PN - Puppet Extended S-Expression Notation
	//
	// A PN forms a directed acyclic graph of nodes. There are four types of nodes:
	//
	// * Literal: A boolean, integer, float, string, or undef
	//
	// * List: An ordered list of nodes
	//
	// * Map: An ordered map of string to node associations
	//
	// * Call: A named list of nodes.
	PN interface {
		// Format produces a compact, Clojure like syntax. Suitable for tests.
		Format(b *bytes.Buffer)

		// ToData produces an object that where all values are of primitive type or
		// slices or maps. This format is suitable for output as JSON or YAML
		//
		ToData() interface{}

		// AsCall turns this PN into an argument in a call or change the name if the PN Is a call already
		AsCall(name string) PN

		// AsParameters returns the PN as a parameter list
		AsParameters() []PN

		// WithName creates a key/value pair from the given name and this PN
		WithName(name string) Entry

		// String returns the Format output as a string
		String() string
	}

	pnError struct {
		message string
	}

	// Entry in hash
	Entry interface {
		Key() string
		Value() PN
	}

	mapEntry struct {
		key   string
		value PN
	}

	ListPN struct {
		elements []PN
	}

	MapPN struct {
		entries []Entry
	}

	LiteralPN struct {
		val interface{}
	}

	CallPN struct {
		ListPN
		name string
	}
)

var keyPattern = regexp.MustCompile(`^[A-Za-z_-][0-9A-Za-z_-]*$`)

// Represent the Reported using Puppet Extended S-Expression Notation (PN)
func ReportedToPN(ri issue.Reported) PN {
	return Map([]Entry{
		Literal(ri.Code()).WithName(`code`),
		Literal(ri.Severity().String()).WithName(`severity`),
		Literal(ri.Error()).WithName(`message`)})
}

func (e *pnError) Error() string {
	return e.message
}

func List(elements []PN) PN {
	return &ListPN{elements}
}

func Map(entries []Entry) PN {
	for _, e := range entries {
		if !keyPattern.MatchString(e.Key()) {
			panic(pnError{fmt.Sprintf("key '%s' does not conform to pattern %s",
				e.Key(), keyPattern.String())})
		}
	}
	return &MapPN{entries}
}

func Literal(val interface{}) PN {
	return &LiteralPN{val}
}

func Call(name string, elements ...PN) PN {
	return &CallPN{ListPN{elements}, name}
}

func ToString(pn PN) string {
	b := bytes.NewBufferString(``)
	pn.Format(b)
	return b.String()
}

func (pn *ListPN) AsCall(name string) PN {
	return Call(name, pn.elements...)
}

func (pn *ListPN) AsParameters() []PN {
	return pn.elements
}

func (pn *ListPN) Format(b *bytes.Buffer) {
	b.WriteByte('[')
	formatElements(pn.elements, b)
	b.WriteByte(']')
}

func (pn *ListPN) ToData() interface{} {
	me := make([]interface{}, len(pn.elements))
	for idx, op := range pn.elements {
		me[idx] = op.ToData()
	}
	return me
}

func (pn *ListPN) String() string {
	return ToString(pn)
}

func (pn *ListPN) WithName(name string) Entry {
	return &mapEntry{name, pn}
}

func (pn *CallPN) AsCall(name string) PN {
	return &CallPN{ListPN{pn.elements}, name}
}

func (pn *CallPN) AsParameters() []PN {
	return pn.elements
}

func (pn *CallPN) Format(b *bytes.Buffer) {
	b.WriteByte('(')
	b.WriteString(pn.name)
	if len(pn.elements) > 0 {
		b.WriteByte(' ')
		formatElements(pn.elements, b)
	}
	b.WriteByte(')')
}

func (pn *CallPN) ToData() interface{} {
	top := len(pn.elements)
	args := make([]interface{}, 0, top+1)
	args = append(args, pn.name)
	if top > 0 {
		params := pn.ListPN.ToData()
		args = append(args, params.([]interface{})...)
	}
	return map[string]interface{}{`^`: args}
}

func (pn *CallPN) String() string {
	return ToString(pn)
}

func (pn *CallPN) WithName(name string) Entry {
	return &mapEntry{name, pn}
}

func (e *mapEntry) Key() string {
	return e.key
}

func (e *mapEntry) Value() PN {
	return e.value
}

func (pn *MapPN) AsCall(name string) PN {
	return Call(name, pn)
}

func (pn *MapPN) AsParameters() []PN {
	return []PN{pn}
}

func (pn *MapPN) Format(b *bytes.Buffer) {
	b.WriteByte('{')
	if top := len(pn.entries); top > 0 {
		formatEntry(pn.entries[0], b)
		for idx := 1; idx < top; idx++ {
			b.WriteByte(' ')
			formatEntry(pn.entries[idx], b)
		}
	}
	b.WriteByte('}')
}

func formatEntry(entry Entry, b *bytes.Buffer) {
	b.WriteByte(':')
	b.WriteString(entry.Key())
	b.WriteByte(' ')
	entry.Value().Format(b)
}

func (pn *MapPN) ToData() interface{} {
	top := len(pn.entries) * 2
	args := make([]interface{}, 0, top)
	for _, entry := range pn.entries {
		args = append(args, entry.Key(), entry.Value().ToData())
	}
	return map[string]interface{}{`#`: args}
}

func (pn *MapPN) String() string {
	return ToString(pn)
}

func (pn *MapPN) WithName(name string) Entry {
	return &mapEntry{name, pn}
}

func (pn *LiteralPN) AsCall(name string) PN {
	return Call(name, pn)
}

func (pn *LiteralPN) AsParameters() []PN {
	return []PN{pn}
}

// Strip zeroes between last significant digit and end or exponent. The
// zero following the decimal point is considered significant.
var StripTrailingZeroes = regexp.MustCompile(`\A(.*(?:\.0|[1-9]))0+(e[+-]?\d+)?\z`)

func (pn *LiteralPN) Format(b *bytes.Buffer) {
	switch pn.val.(type) {
	case nil:
		b.WriteString(`nil`)
	case string:
		DoubleQuote(pn.val.(string), b)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		Fprintf(b, `%d`, pn.val)
	case float32, float64:
		str := fmt.Sprintf(`%.16g`, pn.val)
		// We want 16 digit precision that overflows into scientific notation and no trailing zeroes
		if strings.IndexByte(str, '.') < 0 && strings.IndexByte(str, 'e') < 0 {
			// %g sometimes yields an integer number without decimals or scientific
			// notation. Scientific notation must then be used to retain type information
			str = fmt.Sprintf(`%.16e`, pn.val)
		}

		if groups := StripTrailingZeroes.FindStringSubmatch(str); groups != nil {
			b.WriteString(groups[1])
			b.WriteString(groups[2])
		} else {
			b.WriteString(str)
		}
	case bool:
		Fprintf(b, `%t`, pn.val)
	default:
		Fprintf(b, `%v`, pn.val)
	}
}

func (pn *LiteralPN) ToData() interface{} {
	return pn.val
}

func (pn *LiteralPN) String() string {
	return ToString(pn)
}

func (pn *LiteralPN) WithName(name string) Entry {
	return &mapEntry{name, pn}
}

func formatElements(elements []PN, b *bytes.Buffer) {
	top := len(elements)
	if top > 0 {
		elements[0].Format(b)
		for idx := 1; idx < top; idx++ {
			b.WriteByte(' ')
			elements[idx].Format(b)
		}
	}
}

func Fprintf(w io.Writer, format string, a ...interface{}) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		panic(err)
	}
}

func Fprintln(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		panic(err)
	}
}

func Println(a ...interface{}) {
	_, err := fmt.Println(a...)
	if err != nil {
		panic(err)
	}
}
