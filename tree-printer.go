package main

import (
	"strconv"
	"strings"
)

func TreePrinter() *treePrinter {
	return &treePrinter{
		sb:     &strings.Builder{},
		indent: 0,
	}
}

func (t *treePrinter) String() string {
	return t.sb.String()
}

type treePrinter struct {
	sb     *strings.Builder
	indent int
}

func (t *treePrinter) writeIndent() {
	t.sb.WriteString(strings.Repeat("  ", t.indent))
}

func (t *treePrinter) VisitName(name string) {
	t.writeIndent()
	t.sb.WriteString("name '")
	t.sb.WriteString(name)
	t.sb.WriteString("'\n")
}

func (t *treePrinter) VisitNumber(value float64) {
	t.writeIndent()
	t.sb.WriteString("number ")
	t.sb.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	t.sb.WriteString("\n")
}

func (t *treePrinter) VisitAssign(name string, right Expression) {
	t.writeIndent()
	t.sb.WriteString("assign\n")
	t.indent++
	t.writeIndent()
	t.sb.WriteString("name '")
	t.sb.WriteString(name)
	t.sb.WriteString("'\n")
	right.Visit(t)
	t.indent--
}

func (t *treePrinter) VisitConditional(condition, thenBranch, elseBranch Expression) {
	t.writeIndent()
	t.sb.WriteString("if\n")
	t.indent++
	condition.Visit(t)
	thenBranch.Visit(t)
	elseBranch.Visit(t)
	t.indent--
}

func (t *treePrinter) VisitCall(callee Expression, arguments []Expression) {
	t.writeIndent()
	t.sb.WriteString("call\n")
	t.indent++
	callee.Visit(t)
	for _, arg := range arguments {
		arg.Visit(t)
	}
	t.indent--
}

func (t *treePrinter) VisitPrefix(operator TokenType, right Expression) {
	t.writeIndent()
	t.sb.WriteString("prefix '")
	t.sb.WriteString(string(operator))
	t.sb.WriteString("'\n")
	t.indent++
	right.Visit(t)
	t.indent--
}

func (t *treePrinter) VisitPostfix(left Expression, operator TokenType) {
	t.writeIndent()
	t.sb.WriteString("postfix '")
	t.sb.WriteString(string(operator))
	t.sb.WriteString("'\n")
	t.indent++
	left.Visit(t)
	t.indent--
}

func (t *treePrinter) VisitInfix(left Expression, operator TokenType, right Expression) {
	t.writeIndent()
	t.sb.WriteString("infix '")
	t.sb.WriteString(string(operator))
	t.sb.WriteString("'\n")
	t.indent++
	left.Visit(t)
	right.Visit(t)
	t.indent--
}
