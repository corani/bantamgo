package evaluator

import (
	"errors"
	"log"
	"math"

	"github.com/corani/bantamgo/ast"
	"github.com/corani/bantamgo/lexer"
)

var ErrStackUnderflow = errors.New("stack underflow")

type Function = func([]float64) float64

type SymbolKind int

const (
	SymbolKindNumber SymbolKind = iota
	SymbolKindFunction
	SymbolKindUndefined
)

type Symbol struct {
	Name       string
	Kind       SymbolKind
	AsNumber   float64
	AsFunction Function
}

func New() *eval {
	res := &eval{
		stack:  make([]Symbol, 0),
		locals: make(map[string]Symbol),
	}

	res.defineFunction("pow", func(args []float64) float64 {
		if len(args) != 2 {
			return 0
		}

		return math.Pow(args[0], args[1])
	})

	return res
}

type eval struct {
	stack  []Symbol
	locals map[string]Symbol
}

func (e *eval) Answer() float64 {
	return e.popNumber()
}

func (e *eval) VisitBlock(expressions []ast.Expression) {
	for _, expr := range expressions {
		expr.Visit(e)
	}
}

func (e *eval) VisitName(name string) {
	if val, ok := e.locals[name]; ok {
		e.push(val)
	} else {
		// TODO(daniel): we should probably do this in an earlier type-checking phase.
		e.push(Symbol{Name: name, Kind: SymbolKindUndefined})
	}
}

func (e *eval) VisitNumber(value float64) {
	e.pushNumber(value)
}

func (e *eval) VisitAssign(name string, right ast.Expression) {
	right.Visit(e)

	e.defineNumber(name, e.popNumber())
}

func (e *eval) VisitConditional(condition, thenBranch, elseBranch ast.Expression) {
	condition.Visit(e)

	if int64(e.popNumber()) != 0 {
		thenBranch.Visit(e)
	} else {
		elseBranch.Visit(e)
	}
}

func (e *eval) VisitCall(callee ast.Expression, arguments []ast.Expression) {
	callee.Visit(e)

	fn := e.popFunction()
	args := make([]float64, 0, len(arguments))

	for _, arg := range arguments {
		arg.Visit(e)
		args = append(args, e.popNumber())
	}

	ans := fn(args)

	e.pushNumber(ans)
}

func (e *eval) VisitPrefix(operator lexer.TokenType, right ast.Expression) {
	right.Visit(e)

	val := e.popNumber()

	switch operator {
	case lexer.TypePlus:
		e.pushNumber(val)
	case lexer.TypeMinus:
		e.pushNumber(-val)
	case lexer.TypeTilde:
		e.pushNumber(float64(^int64(val)))
	case lexer.TypeBang:
		if int64(val) == 0 {
			e.pushNumber(1)
		} else {
			e.pushNumber(0)
		}
	}
}

func (e *eval) VisitInfix(left ast.Expression, operator lexer.TokenType, right ast.Expression) {
	left.Visit(e)
	lhs := e.popNumber()

	right.Visit(e)
	rhs := e.popNumber()

	switch operator {
	case lexer.TypePlus:
		e.pushNumber(lhs + rhs)
	case lexer.TypeMinus:
		e.pushNumber(lhs - rhs)
	case lexer.TypeAsterisk:
		e.pushNumber(lhs * rhs)
	case lexer.TypeSlash:
		e.pushNumber(lhs / rhs)
	}
}

func (e *eval) VisitPostfix(left ast.Expression, operator lexer.TokenType) {
	left.Visit(e)

	val := uint64(e.popNumber())

	switch operator {
	case lexer.TypeBang:
		ans := uint64(1)

		for i := uint64(1); i <= val; i++ {
			ans *= i
		}

		e.pushNumber(float64(ans))
	}
}

func (e *eval) define(value Symbol) {
	e.locals[value.Name] = value
}

func (e *eval) defineNumber(name string, value float64) {
	e.define(Symbol{Name: name, Kind: SymbolKindNumber, AsNumber: value})
}

func (e *eval) defineFunction(name string, fn Function) {
	e.define(Symbol{Name: name, Kind: SymbolKindFunction, AsFunction: fn})
}

func (e *eval) push(value Symbol) {
	e.stack = append(e.stack, value)
}

func (e *eval) pushNumber(value float64) {
	e.push(Symbol{Kind: SymbolKindNumber, AsNumber: value})
}

func (e *eval) pop() (*Symbol, error) {
	if len(e.stack) == 0 {
		return nil, ErrStackUnderflow
	}

	value := e.stack[len(e.stack)-1]
	e.stack = e.stack[:len(e.stack)-1]

	return &value, nil
}

func (e *eval) popNumber() float64 {
	// TODO(daniel): not sure if this error recovery is a good idea.
	defaultVal := 0.0

	val, err := e.pop()
	if err != nil {
		log.Println(err)
	} else {
		switch val.Kind {
		case SymbolKindNumber:
			return val.AsNumber
		case SymbolKindUndefined:
			log.Printf("Undefined symbol %q", val.Name)
		default:
			log.Printf("%q is not a number", val.Name)
		}
	}

	return defaultVal
}

func (e *eval) popFunction() Function {
	// TODO(daniel): not sure if this error recovery is a good idea.
	defaultFunc := func([]float64) float64 { return 0 }

	val, err := e.pop()
	if err != nil {
		log.Println(err)
	} else {
		switch val.Kind {
		case SymbolKindFunction:
			return val.AsFunction
		case SymbolKindUndefined:
			log.Printf("Undefined symbol %q", val.Name)
		default:
			log.Printf("%q is not a function", val.Name)
		}
	}

	return defaultFunc
}
