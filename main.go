package main

import (
	"log"
)

func main() {
	input := "a = 1.1 + c * d ^ e - f / g"
	log.Println("input:", input)

	lexer := Lexer(input)
	parser := Parser(lexer)

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
