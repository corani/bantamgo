package main

import "fmt"

type parser struct {
	tokens          *lexer
	read            []Token
	prefixParselets map[TokenType]PrefixParselet
	infixParselets  map[TokenType]InfixParselet
}

func Parser(l *lexer) *parser {
	result := &parser{
		tokens:          l,
		read:            nil,
		prefixParselets: make(map[TokenType]PrefixParselet),
		infixParselets:  make(map[TokenType]InfixParselet),
	}

	// Register special parselets
	result.registerPrefix(TypeName, NameParselet())
	result.registerPrefix(TypeNumber, NumberParselet())
	result.registerPrefix(TypeLParen, GroupParselet())
	result.registerInfix(TypeAssign, AssignParselet())
	result.registerInfix(TypeQuestion, ConditionalParselet())
	result.registerInfix(TypeLParen, CallParselet())

	// Register simple prefix operators
	result.registerPrefix(TypePlus, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(TypeMinus, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(TypeTilde, PrefixOperatorParselet(PrecPrefix))
	result.registerPrefix(TypeBang, PrefixOperatorParselet(PrecPrefix))

	// Register postfix factorial operator
	result.registerPostfix(TypeBang, PostfixOperatorParselet(PrecPostfix))

	// Register left-associative infix operators
	result.registerInfix(TypePlus, InfixOperatorParselet(PrecSum, AssocLeft))
	result.registerInfix(TypeMinus, InfixOperatorParselet(PrecSum, AssocLeft))
	result.registerInfix(TypeAsterisk, InfixOperatorParselet(PrecProduct, AssocLeft))
	result.registerInfix(TypeSlash, InfixOperatorParselet(PrecProduct, AssocLeft))

	// Register right-associative infix operators
	result.registerInfix(TypeCaret, InfixOperatorParselet(PrecExponent, AssocRight))

	return result
}

func (p *parser) registerPrefix(tt TokenType, parselet PrefixParselet) {
	p.prefixParselets[tt] = parselet
}

func (p *parser) registerPostfix(tt TokenType, parselet InfixParselet) {
	p.infixParselets[tt] = parselet
}

func (p *parser) registerInfix(tt TokenType, parselet InfixParselet) {
	p.infixParselets[tt] = parselet
}

func (p *parser) ParseExpression() (Expression, error) {
	return p.parseExpression(0)
}

func (p *parser) parseExpression(precedence Precedence) (Expression, error) {
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

func (p *parser) match(t TokenType) bool {
	if p.lookAhead(0).Type != t {
		return false
	}

	p.consume()

	return true
}

func (p *parser) expect(t TokenType) {
	if !p.match(t) {
		panic("Expected token " + string(t) + " but got " + p.lookAhead(0).Text)
	}
}

func (p *parser) consume() Token {
	p.lookAhead(0)

	result := p.read[0]
	p.read = p.read[1:]

	return result
}

func (p *parser) lookAhead(distance int) Token {
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
