package ast

import (
	"bytes"
	"fmt"
	"strings"
)

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

// BOOLEAN LITERAL
type BooleanLiteral struct {
	Value bool
}

func (bl *BooleanLiteral) expressionNode() {}
func (bl *BooleanLiteral) Literal() string {
	return fmt.Sprintf("%t", bl.Value)
}
func (bl *BooleanLiteral) String() string {
	return fmt.Sprintf("%t", bl.Value)
}
func (bl *BooleanLiteral) DebugString() string {
	return fmt.Sprintf("%t [%T]", bl.Value, bl)
}

// NULL LITERAL
type NullLiteral struct {
}

func (nl *NullLiteral) expressionNode() {}
func (nl *NullLiteral) Literal() string { return "null" }
func (nl *NullLiteral) String() string  { return "null" }
func (nl *NullLiteral) DebugString() string {
	return fmt.Sprintf("null [%T]", nl)
}

// ARRAY LITERAL
type ArrayLiteral struct {
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) Literal() string { return al.String() }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	elems := []string{}
	for _, e := range al.Elements {
		elems = append(elems, e.String())
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}
func (al *ArrayLiteral) DebugString() string {
	var out bytes.Buffer
	out.WriteString("[")
	elems := []string{}
	for _, e := range al.Elements {
		elems = append(elems, e.DebugString())
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	out.WriteString(fmt.Sprintf(" [%T]", al))
	return out.String()
}

// DICT LITERAL
type DictLiteral struct {
	Elements map[string]Expression
}

func (dl *DictLiteral) expressionNode() {}
func (dl *DictLiteral) Literal() string { return dl.String() }
func (dl *DictLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	elems := []string{}
	for k, v := range dl.Elements {
		elems = append(elems, fmt.Sprintf("%s: %s", k, v.String()))
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("}")
	return out.String()
}
func (dl *DictLiteral) DebugString() string {
	var out bytes.Buffer
	out.WriteString("{")
	elems := []string{}
	for k, v := range dl.Elements {
		elems = append(elems, fmt.Sprintf("%s: %s", k, v.DebugString()))
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("}")
	out.WriteString(fmt.Sprintf(" [%T]", dl))
	return out.String()
}
