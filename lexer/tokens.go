package lexer

const (
	TokenEnd = 0

	// Binary ops
	TokenAssign         = 1
	TokenAddAssign      = 2
	TokenSubtractAssign = 3

	TokenMultiply  = 10
	TokenDivide    = 11
	TokenRemainder = 12
	TokenSubtract  = 13
	TokenAdd       = 14

	TokenLshift = 20
	TokenRshift = 21

	TokenEqual        = 30
	TokenNotEqual     = 31
	TokenLess         = 32
	TokenLessEqual    = 33
	TokenGreater      = 34
	TokenGreaterEqual = 35

	TokenMatch    = 40
	TokenNotMatch = 41

	TokenLcollect  = 50
	TokenLlcollect = 51

	TokenRcollect  = 60
	TokenRrcollect = 61

	TokenFarrow = 70
	TokenParrow = 71

	TokenInEdge     = 72
	TokenInEdgeSub  = 73
	TokenOutEdge    = 74
	TokenOutEdgeSub = 75

	// Unary ops
	TokenNot  = 80
	TokenAt   = 81
	TokenAtat = 82

	// ()
	TokenLp   = 90
	TokenWslp = 91
	TokenRp   = 92

	// []
	TokenLb        = 100
	TokenListstart = 101
	TokenRb        = 102

	// {}
	TokenLc   = 110
	TokenSelc = 111
	TokenRc   = 112

	// | |
	TokenPipe    = 120
	TokenPipeEnd = 121

	// EPP
	TokenEppEnd       = 130
	TokenEppEndTrim   = 131
	TokenRenderExpr   = 132
	TokenRenderString = 133

	// Separators
	TokenComma     = 140
	TokenDot       = 141
	TokenQmark     = 142
	TokenColon     = 143
	TokenSemicolon = 144

	// Strings with semantics
	TokenIdentifier         = 150
	TokenString             = 151
	TokenInteger            = 152
	TokenFloat              = 153
	TokenBoolean            = 154
	TokenConcatenatedString = 155
	TokenHeredoc            = 156
	TokenVariable           = 157
	TokenRegexp             = 158
	TokenTypeName           = 159

	// Keywords
	TokenAnd         = 200
	TokenApplication = 201
	TokenAttr        = 202
	TokenCase        = 203
	TokenClass       = 204
	TokenConsumes    = 205
	TokenDefault     = 206
	TokenDefine      = 207
	TokenFunction    = 208
	TokenIf          = 209
	TokenIn          = 210
	TokenInherits    = 211
	TokenElse        = 212
	TokenElsif       = 213
	TokenNode        = 214
	TokenOr          = 215
	TokenPlan        = 216
	TokenPrivate     = 217
	TokenProduces    = 218
	TokenSite        = 219
	TokenType        = 220
	TokenUndef       = 221
	TokenUnless      = 222
)

var TokenMap = map[int]string{
	TokenEnd: `EOF`,

	// Binary ops
	TokenAssign:         `=`,
	TokenAddAssign:      `+=`,
	TokenSubtractAssign: `-=`,

	TokenMultiply:  `*`,
	TokenDivide:    `/`,
	TokenRemainder: `%`,
	TokenSubtract:  `-`,
	TokenAdd:       `+`,

	TokenLshift: `<<`,
	TokenRshift: `>>`,

	TokenEqual:        `==`,
	TokenNotEqual:     `!=`,
	TokenLess:         `<`,
	TokenLessEqual:    `<=`,
	TokenGreater:      `>`,
	TokenGreaterEqual: `>=`,

	TokenMatch:    `=~`,
	TokenNotMatch: `!~`,

	TokenLcollect:  `<|`,
	TokenLlcollect: `<<|`,

	TokenRcollect:  `|>`,
	TokenRrcollect: `|>>`,

	TokenFarrow: `=>`,
	TokenParrow: `+>`,

	TokenInEdge:     `->`,
	TokenInEdgeSub:  `~>`,
	TokenOutEdge:    `<-`,
	TokenOutEdgeSub: `<~`,

	// Unary ops
	TokenNot:  `!`,
	TokenAt:   `@`,
	TokenAtat: `@@`,

	TokenComma: `,`,

	// ()
	TokenLp:   `(`,
	TokenWslp: `(`,
	TokenRp:   `)`,

	// []
	TokenLb:        `[`,
	TokenListstart: `[`,
	TokenRb:        `]`,

	// {}
	TokenLc:   `{`,
	TokenSelc: `{`,
	TokenRc:   `}`,

	// | |
	TokenPipe:    `|`,
	TokenPipeEnd: `|`,

	// EPP
	TokenEppEnd:       `%>`,
	TokenEppEndTrim:   `-%>`,
	TokenRenderExpr:   `<%=`,
	TokenRenderString: `epp text`,

	// Separators
	TokenDot:       `.`,
	TokenQmark:     `?`,
	TokenColon:     `:`,
	TokenSemicolon: `;`,

	// Strings with semantics
	TokenIdentifier:         `identifier`,
	TokenString:             `string literal`,
	TokenInteger:            `integer literal`,
	TokenFloat:              `float literal`,
	TokenBoolean:            `boolean literal`,
	TokenConcatenatedString: `dq string literal`,
	TokenHeredoc:            `heredoc`,
	TokenVariable:           `variable`,
	TokenRegexp:             `regexp`,
	TokenTypeName:           `type name`,

	// Keywords
	TokenAnd:         `and`,
	TokenApplication: `application`,
	TokenAttr:        `attr`,
	TokenCase:        `case`,
	TokenClass:       `class`,
	TokenConsumes:    `consumes`,
	TokenDefault:     `default`,
	TokenDefine:      `define`,
	TokenFunction:    `function`,
	TokenIf:          `if`,
	TokenIn:          `in`,
	TokenInherits:    `inherits`,
	TokenElse:        `else`,
	TokenElsif:       `elsif`,
	TokenNode:        `node`,
	TokenOr:          `or`,
	TokenPlan:        `plan`,
	TokenPrivate:     `private`,
	TokenProduces:    `produces`,
	TokenSite:        `site`,
	TokenType:        `type`,
	TokenUndef:       `undef`,
	TokenUnless:      `unless`,
}

var Keywords = map[string]int{
	TokenMap[TokenApplication]: TokenApplication,
	TokenMap[TokenAnd]:         TokenAnd,
	TokenMap[TokenAttr]:        TokenAttr,
	TokenMap[TokenCase]:        TokenCase,
	TokenMap[TokenClass]:       TokenClass,
	TokenMap[TokenConsumes]:    TokenConsumes,
	TokenMap[TokenDefault]:     TokenDefault,
	TokenMap[TokenDefine]:      TokenDefine,
	`false`:                    TokenBoolean,
	TokenMap[TokenFunction]:    TokenFunction,
	TokenMap[TokenElse]:        TokenElse,
	TokenMap[TokenElsif]:       TokenElsif,
	TokenMap[TokenIf]:          TokenIf,
	TokenMap[TokenIn]:          TokenIn,
	TokenMap[TokenInherits]:    TokenInherits,
	TokenMap[TokenNode]:        TokenNode,
	TokenMap[TokenOr]:          TokenOr,
	TokenMap[TokenPlan]:        TokenPlan,
	TokenMap[TokenPrivate]:     TokenPrivate,
	TokenMap[TokenProduces]:    TokenProduces,
	TokenMap[TokenSite]:        TokenSite,
	`true`:                     TokenBoolean,
	TokenMap[TokenType]:        TokenType,
	TokenMap[TokenUndef]:       TokenUndef,
	TokenMap[TokenUnless]:      TokenUnless,
}

func IsKeywordToken(token int) bool {
	return token >= TokenAnd && token <= TokenUnless
}
