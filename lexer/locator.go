package lexer

import (
	"sort"
	"unicode/utf8"
)

type Locator struct {
	string    string
	file      string
	lineIndex []int
}

func NewLocator(file, content string) *Locator {
	return &Locator{string: content, file: file}
}

func (e *Locator) String() string {
	return e.string
}

func (e *Locator) File() string {
	return e.file
}

// Return the line in the source for the given byte offset
func (e *Locator) LineForOffset(offset int) int {
	return sort.SearchInts(e.getLineIndex(), offset+1)
}

// Return the position on a line in the source for the given byte offset
func (e *Locator) PosOnLine(offset int) int {
	return e.offsetOnLine(offset) + 1
}

func (e *Locator) getLineIndex() []int {
	if e.lineIndex == nil {
		li := append(make([]int, 0, 32), 0)
		rdr := NewStringReader(e.string)
		for c, _ := rdr.Next(); c != 0; c, _ = rdr.Next() {
			if c == '\n' {
				li = append(li, rdr.Pos())
			}
		}
		e.lineIndex = li
	}
	return e.lineIndex
}

func (e *Locator) offsetOnLine(offset int) int {
	li := e.getLineIndex()
	line := sort.SearchInts(li, offset+1)
	lineStart := li[line-1]
	if offset == lineStart {
		return 0
	}
	if offset > len(e.string) {
		offset = len(e.string)
	}
	return utf8.RuneCountInString(e.string[lineStart:offset])
}