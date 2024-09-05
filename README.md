# BantamGo

A Go implementation of the Pratt parser for the Bantam language. Based on the excellent 
[Pratt parsers: Expression parsing made easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)
blog post by Bob Nystrom.

I've taken the liberty to add (floating point) numbers, change the parselets from classes to 
functions and extracted the "Print" function for the Expression using the Visitor pattern.

## Update 1

Added support for parsing blocks of statements so you can do things like: 

```
PI = 3.14159265358979323846;
E = 2.71828182845904523536;
pow(PI, 2) + pow(E, 2);
```

Note: `pow` was added as a built-in function, as it's not (yet) possible to create user-defined
functions.

Moreover, I've added a simple evaluator for the expressions, so you can actually run the code.

Example: 

```bash
$ go run . "PI=3.1415; E=2.7182; pow(PI, 2) + pow(E, 2)"
2024/09/05 17:35:15 input: PI=3.1415; E=2.7182; pow(PI, 2) + pow(E, 2)
2024/09/05 17:35:15 parsed: (PI = 3.1415); (E = 2.7182); (pow(PI, 2) + pow(E, 2))
2024/09/05 17:35:15 s-expr: (block (write 'PI' (number 3.1415)) (write 'E' (number 2.7182)) (+ (call (read 'pow') (read 'PI') (number 2) ) (call (read 'pow') (read 'E') (number 2) )) )
2024/09/05 17:35:15 tree:
block
  assign
    name 'PI'
    number 3.1415
  assign
    name 'E'
    number 2.7182
  infix '+'
    call
      name 'pow'
      name 'PI'
      number 2
    call
      name 'pow'
      name 'E'
      number 2
2024/09/05 17:35:15 answer: 17.25763349
```
