package literal

import (
	"github.com/jpicht/puppet-parser/parser"
)

const notLiteral = `not literal`

func ToLiteral(e parser.Expression) (value interface{}, ok bool) {
	defer func() {
		if err := recover(); err != nil {
			if err == notLiteral {
				ok = false
			} else {
				panic(err)
			}
		}
	}()

	value = toLiteral(e)
	ok = true
	return
}

func toLiteral(e parser.Expression) interface{} {
	switch e := e.(type) {
	case *parser.Program:
		return toLiteral(e.Body())
	case *parser.LiteralList:
		elements := e.Elements()
		result := make([]interface{}, len(elements))
		for idx, elem := range elements {
			result[idx] = toLiteral(elem)
		}
		return result
	case *parser.LiteralHash:
		entries := e.Entries()
		result := make(map[interface{}]interface{}, len(entries))
		for _, entry := range entries {
			kh := entry.(*parser.KeyedEntry)
			result[toLiteral(kh.Key())] = toLiteral(kh.Value())
		}
		return result
	case *parser.ConcatenatedString:
		segments := e.Segments()
		if len(segments) == 1 {
			if ls, ok := segments[0].(*parser.LiteralString); ok {
				return ls.Value()
			}
		}
		panic(notLiteral)
	case *parser.HeredocExpression:
		return toLiteral(e.Text())
	case parser.LiteralValue:
		return e.Value()
	default:
		panic(notLiteral)
	}
}
