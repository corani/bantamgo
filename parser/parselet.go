package parser

import (
	"fmt"

	"github.com/corani/bantamgo/ast"
	"github.com/corani/bantamgo/lexer"
)

type PrefixParselet interface {
	Parse(p *parser, t lexer.Token) (ast.Expression, error)
}

type prefixParseletFunc func(p *parser, t lexer.Token) (ast.Expression, error)

func (p prefixParseletFunc) Parse(parser *parser, t lexer.Token) (ast.Expression, error) {
	return p(parser, t)
}

type InfixParselet interface {
	Parse(p *parser, left ast.Expression, t lexer.Token) (ast.Expression, error)
	Precedence() Precedence
}

type infixParselet struct {
	parse func(p *parser, left ast.Expression, t lexer.Token) (ast.Expression, error)
	prec  Precedence
}

func (i *infixParselet) Parse(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
	return i.parse(parser, left, t)
}

func (i *infixParselet) Precedence() Precedence {
	return i.prec
}

// ----- NAME PARSELET -----

func NameParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t lexer.Token) (ast.Expression, error) {
		return ast.NameExpression(t.Text), nil
	})
}

// ----- NUMBER PARSELET -----

func NumberParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t lexer.Token) (ast.Expression, error) {
		return ast.NumberExpression(t.Text)
	})
}

// ----- ASSIGN PARSELET -----

func AssignParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
			right, err := parser.parseExpression(PrecAssignment - 1)
			if err != nil {
				return nil, err
			}

			if name, ok := left.(*ast.NameExpressionNode); ok {
				return ast.AssignExpression(name.Name, right), nil
			}

			return nil, fmt.Errorf("the left-hand of an assignment must be a name")
		},
		prec: PrecAssignment,
	}
}

// ----- CONDITIONAL PARSELET -----

func ConditionalParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
			thenBranch, err := parser.parseExpression(0)
			if err != nil {
				return nil, err
			}

			parser.expect(lexer.TypeColon)

			elseBranch, err := parser.parseExpression(PrecConditional - 1)
			if err != nil {
				return nil, err
			}

			return ast.ConditionalExpression(left, thenBranch, elseBranch), nil
		},
		prec: PrecConditional,
	}
}

// ----- GROUP PARSELET -----

func GroupParselet() PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t lexer.Token) (ast.Expression, error) {
		expr, err := parser.parseExpression(0)
		if err != nil {
			return nil, err
		}

		parser.expect(lexer.TypeRParen)

		return expr, nil
	})
}

// ----- CALL PARSELET -----

func CallParselet() InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
			var args []ast.Expression

			if !parser.match(lexer.TypeRParen) {
				for {
					arg, err := parser.parseExpression(0)
					if err != nil {
						return nil, err
					}

					args = append(args, arg)

					if !parser.match(lexer.TypeComma) {
						break
					}
				}

				parser.expect(lexer.TypeRParen)
			}

			return ast.CallExpression(left, args), nil
		},
		prec: PrecCall,
	}
}

// ----- PREFIX OPERATOR PARSELET -----

func PrefixOperatorParselet(prec Precedence) PrefixParselet {
	return prefixParseletFunc(func(parser *parser, t lexer.Token) (ast.Expression, error) {
		right, err := parser.parseExpression(prec)
		if err != nil {
			return nil, err
		}

		return ast.PrefixExpression(t.Type, right), nil
	})
}

// ----- POSTFIX OPERATOR PARSELET -----

func PostfixOperatorParselet(prec Precedence) InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
			return ast.PostfixExpression(left, t.Type), nil
		},
		prec: prec,
	}
}

// ----- INFIX OPERATOR PARSELET -----

func InfixOperatorParselet(prec Precedence, assoc Associativity) InfixParselet {
	return &infixParselet{
		parse: func(parser *parser, left ast.Expression, t lexer.Token) (ast.Expression, error) {
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

			return ast.InfixExpression(left, t.Type, right), nil
		},
		prec: prec,
	}
}
