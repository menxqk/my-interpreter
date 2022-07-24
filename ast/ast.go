package ast

type Node interface {
	Literal() string
	String() string
	DebugString() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) Literal() string     { return "" }
func (p *Program) String() string      { return "" }
func (p *Program) DebugString() string { return "" }
