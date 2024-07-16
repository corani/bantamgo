package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()

	tt := []struct {
		in, out string
	}{
		// Function call.
		{"a()", "a()"},
		{"a(b)", "a(b)"},
		{"a(b, c)", "a(b, c)"},
		{"a(b)(c)", "a(b)(c)"},
		{"a(b) + c(d)", "(a(b) + c(d))"},
		{"a(b ? c : d, e + f)", "a((b ? c : d), (e + f))"},

		// Unary precedence.
		{"~!-+a", "(~(!(-(+a))))"},
		{"a!!!", "(((a!)!)!)"},

		// Unary and binary predecence.
		{"-a * b", "((-a) * b)"},
		{"!a + b", "((!a) + b)"},
		{"~a ^ b", "((~a) ^ b)"},
		{"-a!", "(-(a!))"},
		{"!a!", "(!(a!))"},

		// Binary precedence.
		{"a = b + c * d ^ e - f / g", "(a = ((b + (c * (d ^ e))) - (f / g)))"},

		// Binary associativity.
		{"a = b = c", "(a = (b = c))"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a ^ b ^ c", "(a ^ (b ^ c))"},

		// Conditional operator.
		{"a ? b : c ? d : e", "(a ? b : (c ? d : e))"},
		{"a ? b ? c : d : e", "(a ? (b ? c : d) : e)"},
		{"a + b ? c * d : e / f", "((a + b) ? (c * d) : (e / f))"},

		// Grouping.
		{"a + (b + c) + d", "((a + (b + c)) + d)"},
		{"a ^ (b + c)", "(a ^ (b + c))"},
		{"(!a)!", "((!a)!)"},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()

			rq := require.New(t)

			lexer := Lexer(tc.in)
			parser := Parser(lexer)
			printer := Printer()

			expr, err := parser.ParseExpression()
			rq.NoError(err)

			expr.Visit(printer)

			rq.Equal(tc.out, printer.String())
		})
	}
}
