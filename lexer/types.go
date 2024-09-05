package lexer

type TokenType rune

const (
	TypeLParen   TokenType = '('
	TypeRParen   TokenType = ')'
	TypeComma    TokenType = ','
	TypeAssign   TokenType = '='
	TypePlus     TokenType = '+'
	TypeMinus    TokenType = '-'
	TypeAsterisk TokenType = '*'
	TypeSlash    TokenType = '/'
	TypeCaret    TokenType = '^'
	TypeTilde    TokenType = '~'
	TypeBang     TokenType = '!'
	TypeQuestion TokenType = '?'
	TypeColon    TokenType = ':'
	TypeSemi     TokenType = ';'
	TypeEOF      TokenType = -1
	TypeName     TokenType = -2
	TypeNumber   TokenType = -3
)

func TokenTypes() []TokenType {
	return []TokenType{
		TypeLParen,
		TypeRParen,
		TypeComma,
		TypeAssign,
		TypePlus,
		TypeMinus,
		TypeAsterisk,
		TypeSlash,
		TypeCaret,
		TypeTilde,
		TypeBang,
		TypeQuestion,
		TypeColon,
		TypeSemi,
		TypeEOF,
		TypeName,
		TypeNumber,
	}
}

type Token struct {
	Type TokenType
	Text string
}

func NewToken(t TokenType, text ...string) Token {
	if len(text) > 0 {
		return Token{Type: t, Text: text[0]}
	}

	return Token{Type: t, Text: ""}
}
