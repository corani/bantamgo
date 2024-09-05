package main

import (
	"log"

	"github.com/corani/bantamgo/evaluator"
	"github.com/corani/bantamgo/lexer"
	"github.com/corani/bantamgo/parser"
	"github.com/corani/bantamgo/printer"
)

func main() {
	// input := "a = 1.1 + c * d ^ e - f / g"
	// input := "PI*E"
	// input := "pow(2, 4)*PI+4!/undefined()+pow"
	// input := `
	// 	PI = 3.14159265358979323846
	// 	E = 2.71828182845904523536
	// 	PI * E
	// `
	input := `
		PI = 3.14159265358979323846
		E = 2.71828182845904523536
		pow(PI, 2) + pow(E, 2);
	`
	log.Println("input:", input)

	// TODO(daniel): support defining functions, so the built-in `pow` can
	// be replaced with something like `pow = (x, y) => x^y` (syntax tbd).
	// TODO(daniel): support comparison operators ('<=', '>=', '==', '!=', etc.)
	// This will require changes to the lexer to distinguish them from the
	// existing single-character operators (e.g. '=' vs '==' and '!' vs '!=').
	// Alternatively we could use 'eq', 'ne', 'lt', 'le', 'gt', 'ge' or similar
	// and punt it to the parser, though that's probably a bad idea.
	lexer := lexer.New(input)
	parser := parser.New(lexer)

	expr, err := parser.ParseExpression()
	if err != nil {
		log.Fatal(err)
	}

	pprint := printer.Printer()
	expr.Visit(pprint)
	log.Println("parsed:", pprint.String())

	sexpr := printer.SExpr()
	expr.Visit(sexpr)
	log.Println("s-expr:", sexpr.String())

	tree := printer.TreePrinter()
	expr.Visit(tree)
	log.Print("tree:\n" + tree.String())

	// TODO: type-checking

	eval := evaluator.New()
	expr.Visit(eval)
	log.Println("answer:", eval.Answer())

	// TODO: codegen?
}
