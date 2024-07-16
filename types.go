package main

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

type Precedence int

const (
	PrecUnknown     Precedence = 0
	PrecAssignment  Precedence = 1
	PrecConditional Precedence = 2
	PrecSum         Precedence = 3
	PrecProduct     Precedence = 4
	PrecExponent    Precedence = 5
	PrecPrefix      Precedence = 6
	PrecPostfix     Precedence = 7
	PrecCall        Precedence = 8
)

type Associativity bool

const (
	AssocLeft  Associativity = false
	AssocRight Associativity = true
)
