package main

import "strconv"

type Expression interface {
	Visit(v Visitor)
}

// ----- NAME EXPRESSION -----

func NameExpression(name string) *nameExpression {
	return &nameExpression{Name: name}
}

type nameExpression struct {
	Name string
}

func (e *nameExpression) Visit(v Visitor) {
	v.VisitName(e.Name)
}

// ----- NUMBER EXPRESSION -----

func NumberExpression(text string) (*numberExpression, error) {
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil, err
	}

	return &numberExpression{Text: text, Value: val}, nil
}

type numberExpression struct {
	Text  string
	Value float64
}

func (e *numberExpression) Visit(v Visitor) {
	v.VisitNumber(e.Value)
}

// ----- ASSIGN EXPRESSION -----

func AssignExpression(name string, value Expression) *assignExpression {
	return &assignExpression{Name: name, Right: value}
}

type assignExpression struct {
	Name  string
	Right Expression
}

func (e *assignExpression) Visit(v Visitor) {
	v.VisitAssign(e.Name, e.Right)
}

// ----- CONDITIONAL EXPRESSION -----

func ConditionalExpression(condition, thenBranch, elseBranch Expression) *conditionalExpression {
	return &conditionalExpression{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

type conditionalExpression struct {
	Condition  Expression
	ThenBranch Expression
	ElseBranch Expression
}

func (e *conditionalExpression) Visit(v Visitor) {
	v.VisitConditional(e.Condition, e.ThenBranch, e.ElseBranch)
}

// ----- CALL EXPRESSION -----

func CallExpression(callee Expression, args []Expression) *callExpression {
	return &callExpression{Callee: callee, Args: args}
}

type callExpression struct {
	Callee Expression
	Args   []Expression
}

func (e *callExpression) Visit(v Visitor) {
	v.VisitCall(e.Callee, e.Args)
}

// ----- PREFIX EXPRESSION -----

func PrefixExpression(operator TokenType, right Expression) *prefixExpression {
	return &prefixExpression{Operator: operator, Right: right}
}

type prefixExpression struct {
	Operator TokenType
	Right    Expression
}

func (e *prefixExpression) Visit(v Visitor) {
	v.VisitPrefix(e.Operator, e.Right)
}

// ----- POSTFIX EXPRESSION -----

func PostfixExpression(left Expression, operator TokenType) *postfixExpression {
	return &postfixExpression{Operator: operator, Left: left}
}

type postfixExpression struct {
	Operator TokenType
	Left     Expression
}

func (e *postfixExpression) Visit(v Visitor) {
	v.VisitPostfix(e.Left, e.Operator)
}

// ----- INFIX EXPRESSION -----

func InfixExpression(left Expression, operator TokenType, right Expression) *infixExpression {
	return &infixExpression{Left: left, Operator: operator, Right: right}
}

type infixExpression struct {
	Left     Expression
	Operator TokenType
	Right    Expression
}

func (e *infixExpression) Visit(v Visitor) {
	v.VisitInfix(e.Left, e.Operator, e.Right)
}
