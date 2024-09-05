package main

import (
	"log"

	"github.com/corani/bantamgo/lexer"
	"github.com/corani/bantamgo/parser"
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
