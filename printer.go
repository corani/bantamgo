package main

import (
	"strconv"
	"strings"
)

func Printer() *printer {
	return &printer{sb: &strings.Builder{}}
}

func (p *printer) String() string {
	return p.sb.String()
}

type printer struct {
	sb *strings.Builder
}

func (p *printer) VisitName(name string) {
	p.sb.WriteString(name)
}

func (p *printer) VisitNumber(value float64) {
	p.sb.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
}

func (p *printer) VisitAssign(name string, right Expression) {
	p.sb.WriteString("(")
	p.sb.WriteString(name)
	p.sb.WriteString(" = ")
	right.Visit(p)
	p.sb.WriteString(")")
}

func (p *printer) VisitConditional(condition, thenBranch, elseBranch Expression) {
	p.sb.WriteString("(")
	condition.Visit(p)
	p.sb.WriteString(" ? ")
	thenBranch.Visit(p)
	p.sb.WriteString(" : ")
	elseBranch.Visit(p)
	p.sb.WriteString(")")
}

func (p *printer) VisitCall(callee Expression, arguments []Expression) {
	callee.Visit(p)
	p.sb.WriteString("(")
	for i, arg := range arguments {
		if i > 0 {
			p.sb.WriteString(", ")
		}
		arg.Visit(p)
	}
	p.sb.WriteString(")")
}

func (p *printer) VisitPrefix(operator TokenType, right Expression) {
	p.sb.WriteString("(")
	p.sb.WriteString(string(operator))
	right.Visit(p)
	p.sb.WriteString(")")
}

func (p *printer) VisitPostfix(left Expression, operator TokenType) {
	p.sb.WriteString("(")
	left.Visit(p)
	p.sb.WriteString(string(operator))
	p.sb.WriteString(")")
}

func (p *printer) VisitInfix(left Expression, operator TokenType, right Expression) {
	p.sb.WriteString("(")
	left.Visit(p)
	p.sb.WriteString(" ")
	p.sb.WriteString(string(operator))
	p.sb.WriteString(" ")
	right.Visit(p)
	p.sb.WriteString(")")
}
