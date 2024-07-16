package ast

import "github.com/corani/bantamgo/lexer"

type Visitor interface {
	VisitName(name string)
	VisitNumber(value float64)
	VisitAssign(name string, right Expression)
	VisitConditional(condition, thenBranch, elseBranch Expression)
	VisitCall(callee Expression, arguments []Expression)
	VisitPrefix(operator lexer.TokenType, right Expression)
	VisitPostfix(left Expression, operator lexer.TokenType)
	VisitInfix(left Expression, operator lexer.TokenType, right Expression)
}
