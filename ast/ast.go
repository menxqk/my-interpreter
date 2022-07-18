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
