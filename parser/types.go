package parser

type Precedence int

const (
	PrecUnknown     Precedence = 0
	PrecAssignment  Precedence = 1
	PrecConditional Precedence = 2
	PrecSum         Precedence = 3
	PrecProduct     Precedence = 4
	PrecExponent    Precedence = 5
	PrecPrefix      Precedence = 6
	PrecPostfix     Precedence = 7
	PrecCall        Precedence = 8
)

type Associativity bool

const (
	AssocLeft  Associativity = false
	AssocRight Associativity = true
)
