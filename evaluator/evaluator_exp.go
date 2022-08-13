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
	cond := e.Eval(exp.Condition)
	if isError(cond) {
		return cond
	}

	b, ok := cond.(*object.Boolean)
	if !ok {
		return newError("expected %T for boolean, got %T", &object.Boolean{}, cond)
	}

	if b.Value == true && exp.Consequence != nil {
		return e.Eval(exp.Consequence)
	} else if b.Value == false && exp.Alternative != nil {
		return e.Eval(exp.Alternative)
	} else {
		return NULL
	}
}

func (e *Evaluator) evalCallExpression(exp *ast.CallExpression) object.Object {
	fnObj, ok := e.env.Get(exp.Identifier.Name)
	if !ok {
		return newError("%q function not found", exp.Identifier.Name)
	}

	fn, ok := fnObj.(*object.Function)
	if !ok {
		return newError("%q is not a function, got %s", exp.Identifier.Name, fnObj.Type())
	}

	if len(exp.Arguments) != len(fn.Parameters) {
		return newError("wrong number of arguments: %d, expected %d", len(exp.Arguments), len(fn.Parameters))
	}

	params := map[string]object.Object{}
	for i, arg := range exp.Arguments {
		argObj := e.Eval(arg)
		if isError(argObj) {
			return argObj
		}

		param := fn.Parameters[i]
		if argObj.Type() == param.Type {
			// e.env.Set(param.Name, argObj)
			params[param.Name] = argObj
			continue
		} else {
			return newError("wrong type for argument %d, got=%s; expected:%s", i+1, argObj.Type(), param.Type)
		}
	}

	for param, obj := range params {
		e.env.Set(param, obj)
	}

	result := e.Eval(fn.Body)

	for param := range params {
		e.env.Del(param)
	}

	if result.Type() == object.RET_VAL_OBJ {
		resValue := result.(*object.ReturnValue).Value
		if resValue.Type() != fn.Identifier.Type {
			return newError("function %q returned %s, expected %s", exp.Identifier.Name, resValue.Type(), fn.Identifier.Type)
		}
	}

	return result
}

func (e *Evaluator) evalArrayElementExpression(arrElem *ast.ArrayElementExpression) object.Object {
	name := arrElem.Identifier.Name

	obj, ok := e.env.Get(name)
	if !ok {
		return newError("array %q not found", arrElem.Identifier.Name)
	}

	if obj.Type() != object.ARRAY_OBJ {
		return newError("%q not an array", arrElem.Identifier.Name)
	}

	arrObj, ok := obj.(*object.Array)
	if !ok {
		return newError("could not convert %q to array", arrElem.Identifier.Name)
	}

	if arrElem.Index > arrObj.Size-1 {
		return newError("index (%d) out of bounds (%d)", arrElem.Index, arrObj.Size-1)
	}

	if arrElem.Expression != nil {
		newObj := e.Eval(arrElem.Expression)
		if newObj.Type() != arrObj.ArrType {
			return newError("cannot assign %s to %s array", newObj.Type(), arrObj.ArrType)
		}
		arrObj.Elements[arrElem.Index] = newObj
	}

	return arrObj.Elements[arrElem.Index]
}

func (e *Evaluator) evalFunctionExpression(fnExp *ast.FunctionExpression) object.Object {
	return &object.Function{
		Identifier: fnExp.Identifier,
		Parameters: fnExp.Parameters,
		Body:       fnExp.Body,
		Env:        e.env,
	}
}
