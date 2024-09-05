package main

import (
	"log"
	"os"

	"github.com/corani/bantamgo/evaluator"
	"github.com/corani/bantamgo/lexer"
	"github.com/corani/bantamgo/parser"
	"github.com/corani/bantamgo/printer"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: bantamgo <input>")
	}

	input := os.Args[1]

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
