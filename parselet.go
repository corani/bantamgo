package main

import "fmt"

type PrefixParselet interface {
	Parse(p *parser, t Token) (Expression, error)
}

type prefixParseletFunc func(p *parser, t Token) (Expression, error)

func (p prefixParseletFunc) Parse(parser *parser, t Token) (Expression, error) {
	return p(parser, t)
}

type InfixParselet interface {
	Parse(p *parser, left Expression, t Token) (Expression, error)
	Precedence() Precedence
}

type infixParselet struct {
	parse func(p *parser, left Expression, t Token) (Expression, error)
	prec  Precedence
}

func (i *infixParselet) Parse(parser *parser, left Expression, t Token) (Expression, error) {
	return i.parse(parser, left, t)
}

func (i *infixParselet) Precedence() Precedence {
	return i.prec
}

// ----- NAME PARSELET -----

func NameParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t Token) (Expression, error) {
		return NameExpression(t.Text), nil
	})
}

// ----- NUMBER PARSELET -----

func NumberParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t Token) (Expression, error) {
		return NumberExpression(t.Text)
	})
}

// ----- ASSIGN PARSELET -----

func AssignParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left Expression, t Token) (Expression, error) {
			right, err := parser.parseExpression(PrecAssignment - 1)
			if err != nil {
				return nil, err
			}

			if name, ok := left.(*nameExpression); ok {
				return AssignExpression(name.Name, right), nil
			}

			return nil, fmt.Errorf("the left-hand of an assignment must be a name")
		},
		prec: PrecAssignment,
	}
}

// ----- CONDITIONAL PARSELET -----

func ConditionalParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left Expression, t Token) (Expression, error) {
			thenBranch, err := parser.parseExpression(0)
			if err != nil {
				return nil, err
			}

			parser.expect(TypeColon)

			elseBranch, err := parser.parseExpression(PrecConditional - 1)
			if err != nil {
				return nil, err
			}

			return ConditionalExpression(left, thenBranch, elseBranch), nil
		},
		prec: PrecConditional,
	}
}

// ----- GROUP PARSELET -----

func GroupParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t Token) (Expression, error) {
		expr, err := parser.parseExpression(0)
		if err != nil {
			return nil, err
		}

		parser.expect(TypeRParen)

		return expr, nil
	})
}

// ----- CALL PARSELET -----

func CallParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left Expression, t Token) (Expression, error) {
			var args []Expression

			if !parser.match(TypeRParen) {
				for {
					arg, err := parser.parseExpression(0)
					if err != nil {
						return nil, err
					}

					args = append(args, arg)

					if !parser.match(TypeComma) {
						break
					}
				}

				parser.expect(TypeRParen)
			}

			return CallExpression(left, args), nil
		},
		prec: PrecCall,
	}
}

// ----- PREFIX OPERATOR PARSELET -----

func PrefixOperatorParselet(prec Precedence) PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t Token) (Expression, error) {
		right, err := parser.parseExpression(prec)
		if err != nil {
			return nil, err
		}

		return PrefixExpression(t.Type, right), nil
	})
}

// ----- POSTFIX OPERATOR PARSELET -----

func PostfixOperatorParselet(prec Precedence) InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left Expression, t Token) (Expression, error) {
			return PostfixExpression(left, t.Type), nil
		},
		prec: prec,
	}
}

// ----- INFIX OPERATOR PARSELET -----

func InfixOperatorParselet(prec Precedence, assoc Associativity) InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left Expression, t Token) (Expression, error) {
			// To handle right-associative operators like "^", we allow a slightly
			// lower precedence when parsing the right-hand side. This will let a
			// parselet with the same precedence appear on the right, which will then
			// take *this* parselet's result as its left-hand argument.
			prec := prec

			if assoc == AssocRight {
				prec--
			}

			right, err := parser.parseExpression(prec)
			if err != nil {
				return nil, err
			}

			return InfixExpression(left, t.Type, right), nil
		},
		prec: prec,
	}
}
