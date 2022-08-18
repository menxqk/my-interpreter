package evaluator

import (
	"fmt"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

type Evaluator struct {
	env *object.Environment
}

func New() *Evaluator {
	e := &Evaluator{env: object.NewEnvironment()}
	return e
}

func (e *Evaluator) Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return e.evalProgram(node)

	// Statements
	case *ast.ExpressionStatement:
		return e.evalExpressionStatement(node)
	case *ast.BlockStatement:
		return e.evalBlockStatement(node)
	case *ast.FunctionDeclarationStatement:
		return e.evalFunctionDeclarationStatement(node)
	case *ast.ArrayDeclarationStatement:
		return e.evalArrayDeclarationStatement(node)
	case *ast.VariableDeclarationStatement:
		return e.evalVariableDeclarationStatement(node)
	case *ast.AssignmentStatement:
		return e.evalAssignmentStatement(node)
	case *ast.ReturnStatement:
		return e.evalReturnStatement(node)

	// Expressions
	case *ast.Identifier:
		return e.evalIdentifier(node)
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(node)
	case *ast.GroupedExpression:
		return e.evalGroupedExpression(node)
	case *ast.InfixExpression:
		return e.evalInfixExpression(node)
	case *ast.IfExpression:
		return e.evalIfExpression(node)
	case *ast.FunctionExpression:
		return e.evalFunctionExpression(node)
	case *ast.CallExpression:
		return e.evalCallExpression(node)
	case *ast.ArrayElementExpression:
		return e.evalArrayElementExpression(node)
	case *ast.DictElementExpression:
		return e.evalDictElementExpression(node)

	// Literals
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.CharLiteral:
		return &object.Char{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.BooleanLiteral:
		if node.Value == true {
			return TRUE
		}
		return FALSE
	case *ast.NullLiteral:
		return NULL
	case *ast.ArrayLiteral:
		return e.evalArrayLiteral(node)
	case *ast.DictLiteral:
		return e.evalDictLiteral(node)
	default:
		return NULL
	}
}

func (e *Evaluator) evalProgram(program *ast.Program) object.Object {
	var result object.Object = &object.Null{}

	for _, stmt := range program.Statements {
		result = e.Eval(stmt)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
