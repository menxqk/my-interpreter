package ast

import "fmt"

// INTEGER LITERAL
type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) expressionNode()     {}
func (il *IntegerLiteral) Literal() string     { return fmt.Sprintf("%d", il.Value) }
func (il *IntegerLiteral) String() string      { return fmt.Sprintf("%d", il.Value) }
func (il *IntegerLiteral) DebugString() string { return fmt.Sprintf("%d [%T]", il.Value, il) }

// FLOAT LITERAL
type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) expressionNode()     {}
func (fl *FloatLiteral) Literal() string     { return fmt.Sprintf("%.6f", fl.Value) }
func (fl *FloatLiteral) String() string      { return fmt.Sprintf("%.6f", fl.Value) }
func (fl *FloatLiteral) DebugString() string { return fmt.Sprintf("%.6f [%T]", fl.Value, fl) }

// CHAR LITERAL
type CharLiteral struct {
	Value rune
}

func (cl *CharLiteral) expressionNode()     {}
func (cl *CharLiteral) Literal() string     { return fmt.Sprintf("%s", string(cl.Value)) }
func (cl *CharLiteral) String() string      { return fmt.Sprintf("'%s'", string(cl.Value)) }
func (cl *CharLiteral) DebugString() string { return fmt.Sprintf("'%s' [%T]", string(cl.Value), cl) }

// STRING LITERAL
type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode()     {}
func (sl *StringLiteral) Literal() string     { return sl.Value }
func (sl *StringLiteral) String() string      { return fmt.Sprintf("\"%s\"", sl.Value) }
func (sl *StringLiteral) DebugString() string { return fmt.Sprintf("\"%s\" [%T]", sl.Value, sl) }
