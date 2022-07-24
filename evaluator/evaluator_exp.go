package evaluator

import (
	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/object"
)

func (e *Evaluator) evalIdentifier(ident *ast.Identifier) object.Object {
	obj, ok := e.env.Get(ident.Name)
	if !ok {
		return &object.Null{}
	}
	return obj
}

func (e *Evaluator) evalPrefixExpression(exp *ast.PrefixExpression) object.Object {
	// TODO
	return &object.Null{}
}

func (e *Evaluator) evalGroupedExpression(exp *ast.GroupedExpression) object.Object {
	// TODO
	return &object.Null{}
}

func (e *Evaluator) evalInfixExpression(exp *ast.InfixExpression) object.Object {
	// TODO
	return &object.Null{}
}

func (e *Evaluator) evalIfExpression(exp *ast.IfExpression) object.Object {
	// TODO
	return &object.Null{}
}

func (e *Evaluator) evalCallExpression(exp *ast.CallExpression) object.Object {
	// TODO
	return &object.Null{}
}

func (e *Evaluator) evalFunctionExpression(fnExp *ast.FunctionExpression) object.Object {
	return &object.Function{
		Identifier: fnExp.Identifier,
		Parameters: fnExp.Parameters,
		Body:       fnExp.Body,
		Env:        e.env,
	}
}
