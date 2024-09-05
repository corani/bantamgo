package ast

import (
	"strconv"

	"github.com/corani/bantamgo/lexer"
)

type Expression interface {
	Visit(v Visitor)
}

// ----- BLOCK EXPRESSION -----

func BlockExpression(expressions []Expression) *BlockExpressionNode {
	return &BlockExpressionNode{Expressions: expressions}
}

type BlockExpressionNode struct {
	Expressions []Expression
}

func (e *BlockExpressionNode) Visit(v Visitor) {
	v.VisitBlock(e.Expressions)
}

// ----- NAME EXPRESSION -----

func NameExpression(name string) *NameExpressionNode {
	return &NameExpressionNode{Name: name}
}

type NameExpressionNode struct {
	Name string
}

func (e *NameExpressionNode) Visit(v Visitor) {
	v.VisitName(e.Name)
}

// ----- NUMBER EXPRESSION -----

func NumberExpression(text string) (*NumberExpressionNode, error) {
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil, err
	}

	return &NumberExpressionNode{Text: text, Value: val}, nil
}

type NumberExpressionNode struct {
	Text  string
	Value float64
}

func (e *NumberExpressionNode) Visit(v Visitor) {
	v.VisitNumber(e.Value)
}

// ----- ASSIGN EXPRESSION -----

func AssignExpression(name string, value Expression) *AssignExpressionNode {
	return &AssignExpressionNode{Name: name, Right: value}
}

type AssignExpressionNode struct {
	Name  string
	Right Expression
}

func (e *AssignExpressionNode) Visit(v Visitor) {
	v.VisitAssign(e.Name, e.Right)
}

// ----- CONDITIONAL EXPRESSION -----

func ConditionalExpression(condition, thenBranch, elseBranch Expression) *ConditionalExpressionNode {
	return &ConditionalExpressionNode{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

type ConditionalExpressionNode struct {
	Condition  Expression
	ThenBranch Expression
	ElseBranch Expression
}

func (e *ConditionalExpressionNode) Visit(v Visitor) {
	v.VisitConditional(e.Condition, e.ThenBranch, e.ElseBranch)
}

// ----- CALL EXPRESSION -----

func CallExpression(callee Expression, args []Expression) *CallExpressionNode {
	return &CallExpressionNode{Callee: callee, Args: args}
}

type CallExpressionNode struct {
	Callee Expression
	Args   []Expression
}

func (e *CallExpressionNode) Visit(v Visitor) {
	v.VisitCall(e.Callee, e.Args)
}

// ----- PREFIX EXPRESSION -----

func PrefixExpression(operator lexer.TokenType, right Expression) *PrefixExpressionNode {
	return &PrefixExpressionNode{Operator: operator, Right: right}
}

type PrefixExpressionNode struct {
	Operator lexer.TokenType
	Right    Expression
}

func (e *PrefixExpressionNode) Visit(v Visitor) {
	v.VisitPrefix(e.Operator, e.Right)
}

// ----- POSTFIX EXPRESSION -----

func PostfixExpression(left Expression, operator lexer.TokenType) *PostfixExpressionNode {
	return &PostfixExpressionNode{Operator: operator, Left: left}
}

type PostfixExpressionNode struct {
	Operator lexer.TokenType
	Left     Expression
}

func (e *PostfixExpressionNode) Visit(v Visitor) {
	v.VisitPostfix(e.Left, e.Operator)
}

// ----- INFIX EXPRESSION -----

func InfixExpression(left Expression, operator lexer.TokenType, right Expression) *InfixExpressionNode {
	return &InfixExpressionNode{Left: left, Operator: operator, Right: right}
}

type InfixExpressionNode struct {
	Left     Expression
	Operator lexer.TokenType
	Right    Expression
}

func (e *InfixExpressionNode) Visit(v Visitor) {
	v.VisitInfix(e.Left, e.Operator, e.Right)
}
