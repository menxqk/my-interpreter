package ast

import (
	"bytes"
	"fmt"
	"strings"
)

// IDENTIFIER
type Identifier struct {
	Name        string
	Type        string
	TypeLiteral string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Literal() string { return i.Name }
func (i *Identifier) String() string {
	if i.TypeLiteral != "" {
		return fmt.Sprintf("%s %s", i.TypeLiteral, i.Name)
	}
	return i.Name
}
func (i *Identifier) DebugString() string {
	if i.TypeLiteral != "" {
		return fmt.Sprintf("%s %s [%T]", i.TypeLiteral, i.Name, i)
	}
	return fmt.Sprintf("%s [%T]", i.Name, i)
}

// PREFIX EXPRESSION
type PrefixExpression struct {
	Operator   string
	Expression Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) Literal() string { return pe.Operator }
func (pe *PrefixExpression) String() string {
	if pe.Expression != nil {
		return fmt.Sprintf("%s%s", pe.Operator, pe.Expression.String())
	}
	return ""
}
func (pe *PrefixExpression) DebugString() string {
	if pe.Expression != nil {
		return fmt.Sprintf("%s%s [%T]", pe.Operator, pe.Expression.DebugString(), pe)
	}
	return ""
}

// GROUPED EXPRESSION
type GroupedExpression struct {
	Expression Expression
}

func (ge *GroupedExpression) expressionNode() {}
func (ge *GroupedExpression) Literal() string { return "(" }
func (ge *GroupedExpression) String() string {
	if ge.Expression != nil {
		return fmt.Sprintf("(%s)", ge.Expression.String())
	}
	return ""
}
func (ge *GroupedExpression) DebugString() string {
	if ge.Expression != nil {
		return fmt.Sprintf("(%s) [%T]", ge.Expression.DebugString(), ge)
	}
	return ""
}

// INFIX EXPRESSION
type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) Literal() string { return "()" }
func (ie *InfixExpression) String() string {
	if ie.Left != nil && ie.Right != nil {
		return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
	}
	return ""
}
func (ie *InfixExpression) DebugString() string {
	if ie.Left != nil && ie.Right != nil {
		return fmt.Sprintf("(%s %s %s) [%T]", ie.Left.DebugString(), ie.Operator, ie.Right.DebugString(), ie)
	}
	return ""
}

// IF EXPRESSION
type IfExpression struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) Literal() string { return "if" }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	if ie.Condition != nil {
		out.WriteString(fmt.Sprintf("if %s ", ie.Condition.String()))
	}
	if ie.Consequence != nil {
		out.WriteString(fmt.Sprintf("%s ", ie.Consequence.String()))
	}
	if ie.Alternative != nil {
		out.WriteString(fmt.Sprintf("else %s", ie.Alternative.String()))
	}
	return out.String()
}
func (ie *IfExpression) DebugString() string {
	var out bytes.Buffer
	if ie.Condition != nil {
		out.WriteString(fmt.Sprintf("if %s ", ie.Condition.DebugString()))
	}
	if ie.Consequence != nil {
		out.WriteString(fmt.Sprintf("%s", ie.Consequence.DebugString()))
	}
	if ie.Alternative != nil {
		out.WriteString(fmt.Sprintf("else %s", ie.Alternative.DebugString()))
	}
	out.WriteString(fmt.Sprintf(" [%T]", ie))
	return out.String()
}

// FUNCTION EXPRESSION
type FunctionExpression struct {
	Identifier Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fe *FunctionExpression) expressionNode() {}
func (fe *FunctionExpression) Literal() string { return "func" }
func (fe *FunctionExpression) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s %s(", fe.Identifier.TypeLiteral, fe.Identifier.Name))
	params := []string{}
	for _, param := range fe.Parameters {
		params = append(params, fmt.Sprintf("%s %s", param.TypeLiteral, param.Name))
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if fe.Body != nil {
		out.WriteString(fmt.Sprintf(" %s", fe.Body.String()))
	}
	return out.String()
}
func (fe *FunctionExpression) DebugString() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s %s [%T](", fe.Identifier.TypeLiteral, fe.Identifier.Name, fe.Identifier))
	params := []string{}
	for _, param := range fe.Parameters {
		params = append(params, fmt.Sprintf("%s %s [%T]", param.TypeLiteral, param.Name, param))
	}
	out.WriteString(fmt.Sprintf("%s", strings.Join(params, ", ")))
	out.WriteString(")")
	if fe.Body != nil {
		out.WriteString(fmt.Sprintf(" %s", fe.Body.DebugString()))
	}
	out.WriteString(fmt.Sprintf(" [%T]", fe))
	return out.String()
}

// CALL EXPRESSION
type CallExpression struct {
	Identifier Identifier
	Arguments  []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) Literal() string { return "call" }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s(", ce.Identifier.String()))
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
func (ce *CallExpression) DebugString() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s(", ce.Identifier.DebugString()))
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.DebugString())
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	out.WriteString(fmt.Sprintf(" [%T]", ce))
	return out.String()
}

// ARRAY ELEMENT EXPRESSION
type ArrayElementExpression struct {
	Identifier Identifier
	Index      int
	Expression Expression
}

func (aee *ArrayElementExpression) expressionNode() {}
func (aee *ArrayElementExpression) Literal() string { return "element" }
func (aee *ArrayElementExpression) String() string {
	if aee.Expression != nil {
		return fmt.Sprintf("%s[%d] = %s", aee.Identifier.String(), aee.Index, aee.Expression.String())
	}
	return fmt.Sprintf("%s[%d]", aee.Identifier.String(), aee.Index)
}
func (aee *ArrayElementExpression) DebugString() string {
	if aee.Expression != nil {
		return fmt.Sprintf("%s[%d] = %s [%T]", aee.Identifier.String(), aee.Index, aee.Expression.DebugString(), aee)
	}
	return fmt.Sprintf("%s[%d] [%T]", aee.Identifier.String(), aee.Index, aee)
}
