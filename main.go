package main

import (
	"log"

	"github.com/corani/bantamgo/lexer"
	"github.com/corani/bantamgo/parser"
)

func main() {
	input := "a = 1.1 + c * d ^ e - f / g"
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
}
