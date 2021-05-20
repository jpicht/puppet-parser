package lexer

type Lexer interface {
	CurrentToken() int

	NextToken() int

	TokenStartPos() int

	TokenValue() interface{}

	TokenString() string

	AssertToken(token int)

	Locator() *Locator
}
