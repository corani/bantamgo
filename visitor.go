package main

type Visitor interface {
	VisitName(name string)
	VisitNumber(value float64)
	VisitAssign(name string, right Expression)
	VisitConditional(condition, thenBranch, elseBranch Expression)
	VisitCall(callee Expression, arguments []Expression)
	VisitPrefix(operator TokenType, right Expression)
	VisitPostfix(left Expression, operator TokenType)
	VisitInfix(left Expression, operator TokenType, right Expression)
}
