package ast

import "fmt"

// EXPRESSION STATEMENT
type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) statementNode()  {}
func (es *ExpressionStatement) Literal() string { return "EXP_STMT" }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return fmt.Sprintf("%s;", es.Expression.String())
	}
	return ""
}

func (es *ExpressionStatement) DebugString() string {
	if es.Expression != nil {
		return fmt.Sprintf("%s; [%T]", es.Expression.DebugString(), es)
	}
	return ""
}
