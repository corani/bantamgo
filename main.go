package main

import (
	"log"

	"github.com/corani/bantamgo/lexer"
	"github.com/corani/bantamgo/parser"
)

func main() {
	// input := "a = 1.1 + c * d ^ e - f / g"
	// input := "PI*E"
	input := "pow(2, 4)*PI+4!/undefined()+pow"
	log.Println("input:", input)

	// TODO(daniel): support multiple statements, so you can do something like:
	// a = 3;
	// b = 4;
	// pow(a, 2) + pow(b, 2);
	//
	// ==> 25
	lexer := lexer.New(input)
	parser := parser.New(lexer)

	expr, err := parser.ParseExpression()
	if err != nil {
		log.Fatal(err)
	}

	printer := Printer()
	expr.Visit(printer)
	log.Println("parsed:", printer.String())

	sexpr := SExpr()
	expr.Visit(sexpr)
	log.Println("s-expr:", sexpr.String())

	tree := TreePrinter()
	expr.Visit(tree)
	log.Print("tree:\n" + tree.String())

	// TODO: type-checking

	eval := Eval()
	expr.Visit(eval)
	log.Println("answer:", eval.Answer())

	// TODO: codegen?
}
