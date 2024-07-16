# BantamGo

A Go implementation of the Pratt parser for the Bantam language. Based on the excellent 
[Pratt parsers: Expression parsing made easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)
blog post by Bob Nystrom.

I've taken the liberty to add (floating point) numbers, change the parselets from classes to 
functions and extracted the "Print" function for the Expression using the Visitor pattern.
