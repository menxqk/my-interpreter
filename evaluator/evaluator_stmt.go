package evaluator

import (
	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/object"
)

func (e *Evaluator) evalExpressionStatement(stmt *ast.ExpressionStatement) object.Object {
	return e.Eval(stmt.Expression)
}

func (e *Evaluator) evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
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

func (e *Evaluator) evalVariableDeclarationStatement(stmt *ast.VariableDeclarationStatement) object.Object {
	var result object.Object

	name := stmt.Identifier.Name
	varType := stmt.Identifier.Type

	obj := e.Eval(stmt.Expression)

	if varType != obj.Type() {
		return newError("cannot assign %s to %s", obj.Type(), varType)
	}

	result = e.env.Set(name, obj)

	return result
}

func (e *Evaluator) evalFunctionDeclarationStatement(stmt *ast.FunctionDeclarationStatement) object.Object {
	var result object.Object

	obj := e.Eval(stmt.Function)
	fn, ok := obj.(*object.Function)
	if !ok {
		return newError("could not eval function")
	}

	result = e.env.Set(fn.Identifier.Name, fn)

	return result
}

func (e *Evaluator) evalAssignmentStatement(stmt *ast.AssignmentStatement) object.Object {
	var result object.Object

	obj, ok := e.env.Get(stmt.Identifier.Name)
	if !ok {
		return newError("%q not declared", stmt.Identifier.Name)
	}

	expObj := e.Eval(stmt.Expression)

	if expObj.Type() != obj.Type() {
		return newError("cannot assign %s to %s", expObj.Type(), obj.Type())
	}

	result = e.env.Set(stmt.Identifier.Name, expObj)

	return result
}

func (e *Evaluator) evalReturnStatement(stmt *ast.ReturnStatement) object.Object {
	return e.Eval(stmt.ReturnValue)
}
