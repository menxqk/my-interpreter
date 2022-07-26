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
	right := e.Eval(exp.Expression)
	if isError(right) {
		return right
	}

	operator := exp.Operator

	switch operator {
	case "!":
		return e.evalBangOperatorExpression(right)
	case "-":
		return e.evalMinusOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func (e *Evaluator) evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func (e *Evaluator) evalMinusOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.INT_OBJ:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	case object.FLOAT_OBJ:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

func (e *Evaluator) evalGroupedExpression(exp *ast.GroupedExpression) object.Object {
	return e.Eval(exp.Expression)
}

func (e *Evaluator) evalInfixExpression(exp *ast.InfixExpression) object.Object {
	left := e.Eval(exp.Left)
	if isError(left) {
		return left
	}

	right := e.Eval(exp.Right)
	if isError(right) {
		return right
	}

	operator := exp.Operator

	switch operator {
	case "+":
		return e.evalPlusOperator(left, right)
	case "-":
		return e.evalMinusOperator(left, right)
	case "*":
		return e.evalAsteriskOperator(left, right)
	case "/":
		return e.evalSlashOperator(left, right)
	case "==":
		return e.evalEqualOperator(left, right)
	case "!=":
		return e.evalNotEqualOperator(left, right)
	case ">":
		return e.evalGreaterOperator(left, right)
	case ">=":
		return e.evalGreaterEqualOperator(left, right)
	case "<":
		return e.evalLesserOperator(left, right)
	case "<=":
		return e.evalLesserEqualOperator(left, right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
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
