package ast

import (
	"bytes"
	"fmt"
)

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
		return fmt.Sprintf("%s [%T];", es.Expression.DebugString(), es)
	}
	return ""
}

// BLOCK STATEMENT
type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statementNode()  {}
func (bs *BlockStatement) Literal() string { return "BLOCK_STMT" }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, stmt := range bs.Statements {
		out.WriteString(fmt.Sprintf(" %s", stmt.String()))
	}
	out.WriteString(" }")
	return out.String()
}
func (bs *BlockStatement) DebugString() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, stmt := range bs.Statements {
		out.WriteString(fmt.Sprintf(" %s", stmt.DebugString()))
	}
	out.WriteString(" }")
	out.WriteString(fmt.Sprintf(" [%T]", bs))
	return out.String()
}

// VARIABLE DECLARATION STATEMENT
type VariableDeclarationStatement struct {
	Identifier Identifier
	Expression Expression
}

func (vds *VariableDeclarationStatement) statementNode()  {}
func (vds *VariableDeclarationStatement) Literal() string { return "VD_STMT" }
func (vds *VariableDeclarationStatement) String() string {
	if vds.Expression != nil {
		return fmt.Sprintf("%s = %s;", vds.Identifier.String(), vds.Expression.String())
	}
	return fmt.Sprintf("%s;", vds.Identifier.String())
}
func (vds *VariableDeclarationStatement) DebugString() string {
	if vds.Expression != nil {
		return fmt.Sprintf("%s = %s [%T];", vds.Identifier.DebugString(), vds.Expression.DebugString(), vds)
	}
	return fmt.Sprintf("%s [%T];", vds.Identifier.DebugString(), vds)
}

// FUNCTION DECLARATION STATEMENT
type FunctionDeclarationStatement struct {
	Function Expression
}

func (fds *FunctionDeclarationStatement) statementNode()  {}
func (fds *FunctionDeclarationStatement) Literal() string { return "FD_STMT" }
func (fds *FunctionDeclarationStatement) String() string {
	if fds.Function != nil {
		return fmt.Sprintf("%s", fds.Function.String())
	}
	return ""
}
func (fds *FunctionDeclarationStatement) DebugString() string {
	if fds.Function != nil {
		return fmt.Sprintf("%s [%T]", fds.Function.DebugString(), fds)
	}
	return ""
}

// ARRAY DECLARATION STATEMENT
type ArrayDeclarationStatement struct {
	Identifier Identifier
	Size       int
	Expression Expression
}

func (ads *ArrayDeclarationStatement) statementNode()  {}
func (ads *ArrayDeclarationStatement) Literal() string { return "AD_STMT" }
func (ads *ArrayDeclarationStatement) String() string {
	if ads.Expression != nil {
		return fmt.Sprintf("%s[%d] = %s;", ads.Identifier.String(), ads.Size, ads.Expression.String())
	}
	return fmt.Sprintf("%s[%d];", ads.Identifier.String(), ads.Size)
}
func (ads *ArrayDeclarationStatement) DebugString() string {
	if ads.Expression != nil {
		return fmt.Sprintf("%s[%d] = %s [%T];", ads.Identifier.DebugString(), ads.Size, ads.Expression.DebugString(), ads)
	}
	return fmt.Sprintf("%s[%d] [%T];", ads.Identifier.String(), ads.Size, ads)
}

// ASSIGNMENT STATEMENT
type AssignmentStatement struct {
	Identifier Identifier
	Expression Expression
}

func (as *AssignmentStatement) statementNode()  {}
func (as *AssignmentStatement) Literal() string { return "A_STMT" }
func (as *AssignmentStatement) String() string {
	if as.Expression != nil {
		return fmt.Sprintf("%s = %s;", as.Identifier.String(), as.Expression.String())
	}
	return ""
}
func (as *AssignmentStatement) DebugString() string {
	if as.Expression != nil {
		return fmt.Sprintf("%s = %s [%T];", as.Identifier.DebugString(), as.Expression.DebugString(), as)
	}
	return ""
}

// RETURN STATEMENT
type ReturnStatement struct {
	ReturnValue Expression
}

func (re *ReturnStatement) statementNode()  {}
func (re *ReturnStatement) Literal() string { return "return" }
func (re *ReturnStatement) String() string {
	if re.ReturnValue != nil {
		return fmt.Sprintf("return %s;", re.ReturnValue.String())
	}
	return ""
}
func (re *ReturnStatement) DebugString() string {
	if re.ReturnValue != nil {
		return fmt.Sprintf("return %s [%T];", re.ReturnValue.DebugString(), re)
	}
	return ""
}
