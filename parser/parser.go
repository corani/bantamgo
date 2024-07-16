package parser

import (
	"fmt"

	"github.com/corani/bantamgo/ast"
	"github.com/corani/bantamgo/lexer"
)

type parser struct {
	tokens          *lexer.Lexer
	read            []lexer.Token
	prefixParselets map[lexer.TokenType]PrefixParselet
	infixParselets  map[lexer.TokenType]InfixParselet
}

func New(l *lexer.Lexer) *parser {
	result := &parser{
		tokens:          l,
		read:            nil,
		prefixParselets: make(map[lexer.TokenType]PrefixParselet),
		infixParselets:  make(map[lexer.TokenType]InfixParselet),
	}

	// Register special parselets
	result.registerPrefix(lexer.TypeName, NameParselet())
	result.registerPrefix(lexer.TypeNumber, NumberParselet())
	result.registerPrefix(lexer.TypeLParen, GroupParselet())
	result.registerInfix(lexer.TypeAssign, AssignParselet())
	result.registerInfix(lexer.TypeQuestion, ConditionalParselet())
	result.registerInfix(lexer.TypeLParen, CallParselet())

	// Register simple prefix operators
	result.registerPrefix(lexer.TypePlus, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(lexer.TypeMinus, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(lexer.TypeTilde, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(lexer.TypeBang, PrefixOperatorParselet(PrecPrefix))

	// Register postfix factorial operator
	result.registerPostfix(lexer.TypeBang, PostfixOperatorParselet(PrecPostfix))

	// Register left-associative infix operators
	result.registerInfix(lexer.TypePlus, InfixOperatorParselet(PrecSum, AssocLeft))
	result.registerInfix(lexer.TypeMinus, InfixOperatorParselet(PrecSum, AssocLeft))
	result.registerInfix(lexer.TypeAsterisk, InfixOperatorParselet(PrecProduct, AssocLeft))
	result.registerInfix(lexer.TypeSlash, InfixOperatorParselet(PrecProduct, AssocLeft))

	// Register right-associative infix operators
	result.registerInfix(lexer.TypeCaret, InfixOperatorParselet(PrecExponent, AssocRight))

	return result
}

func (p *parser) registerPrefix(tt lexer.TokenType, parselet PrefixParselet) {
	p.prefixParselets[tt] = parselet
}

func (p *parser) registerPostfix(tt lexer.TokenType, parselet InfixParselet) {
	p.infixParselets[tt] = parselet
}

func (p *parser) registerInfix(tt lexer.TokenType, parselet InfixParselet) {
	p.infixParselets[tt] = parselet
}

func (p *parser) ParseExpression() (ast.Expression, error) {
	return p.parseExpression(0)
}

func (p *parser) parseExpression(precedence Precedence) (ast.Expression, error) {
	t := p.consume()

	if prefix, ok := p.prefixParselets[t.Type]; ok {
		left, err := prefix.Parse(p, t)
		if err != nil {
			return nil, err
		}

		for precedence < p.getPrecedence() {
			t = p.consume()

			if infix, ok := p.infixParselets[t.Type]; ok {
				left, err = infix.Parse(p, left, t)
				if err != nil {
					return nil, err
				}
			}
		}

		return left, nil
	}

	return nil, fmt.Errorf("Unexpected token: " + t.Text)
}

func (p *parser) match(t lexer.TokenType) bool {
	if p.lookAhead(0).Type != t {
		return false
	}

	p.consume()

	return true
}

func (p *parser) expect(t lexer.TokenType) {
	if !p.match(t) {
		panic("Expected token " + string(t) + " but got " + p.lookAhead(0).Text)
	}
}

func (p *parser) consume() lexer.Token {
	p.lookAhead(0)

	result := p.read[0]
	p.read = p.read[1:]

	return result
}

func (p *parser) lookAhead(distance int) lexer.Token {
	for distance >= len(p.read) {
		p.read = append(p.read, p.tokens.Next())
	}

	return p.read[distance]
}

func (p *parser) getPrecedence() Precedence {
	if parselet, ok := p.infixParselets[p.lookAhead(0).Type]; ok {
		return parselet.Precedence()
	}

	return PrecUnknown
}
