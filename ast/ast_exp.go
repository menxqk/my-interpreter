package ast

import "fmt"

// IDENTIFIER
type Identifier struct {
	Name     string
	Type     string
	DataType string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Literal() string { return i.Name }
func (i *Identifier) String() string {
	if i.Type != "" {
		return fmt.Sprintf("%s %s", i.Type, i.Name)
	}
	return i.Name
}
func (i *Identifier) DebugString() string {
	if i.Type != "" {
		return fmt.Sprintf("%s %s [%T]", i.Type, i.Name, i)
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
