package main

import (
	"strconv"
	"strings"
)

func SExpr() *sExpr {
	return &sExpr{sb: &strings.Builder{}}
}

func (s *sExpr) String() string {
	return s.sb.String()
}

type sExpr struct {
	sb *strings.Builder
}

func (s *sExpr) VisitName(name string) {
	s.sb.WriteString("(read '")
	s.sb.WriteString(name)
	s.sb.WriteString("')")
}

func (s *sExpr) VisitNumber(value float64) {
	s.sb.WriteString("(number ")
	s.sb.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	s.sb.WriteString(")")
}

func (s *sExpr) VisitAssign(name string, right Expression) {
	s.sb.WriteString("(write '")
	s.sb.WriteString(name)
	s.sb.WriteString("' ")
	right.Visit(s)
	s.sb.WriteString(")")
}

func (s *sExpr) VisitConditional(condition, thenBranch, elseBranch Expression) {
	s.sb.WriteString("(if ")
	condition.Visit(s)
	s.sb.WriteString(" ")
	thenBranch.Visit(s)
	s.sb.WriteString(" ")
	elseBranch.Visit(s)
	s.sb.WriteString(")")
}

func (s *sExpr) VisitCall(callee Expression, arguments []Expression) {
	s.sb.WriteString("(call ")
	callee.Visit(s)
	s.sb.WriteString(" ")
	for _, arg := range arguments {
		arg.Visit(s)
		s.sb.WriteString(" ")
	}
	s.sb.WriteString(")")
}

func (s *sExpr) VisitPrefix(operator TokenType, right Expression) {
	s.sb.WriteString("(prefix")
	s.sb.WriteString(string(operator))
	s.sb.WriteString(" ")
	right.Visit(s)
	s.sb.WriteString(")")
}

func (s *sExpr) VisitPostfix(left Expression, operator TokenType) {
	s.sb.WriteString("(postfix")
	s.sb.WriteString(string(operator))
	s.sb.WriteString(" ")
	left.Visit(s)
	s.sb.WriteString(")")
}

func (s *sExpr) VisitInfix(left Expression, operator TokenType, right Expression) {
	s.sb.WriteString("(")
	s.sb.WriteString(string(operator))
	s.sb.WriteString(" ")
	left.Visit(s)
	s.sb.WriteString(" ")
	right.Visit(s)
	s.sb.WriteString(")")
}