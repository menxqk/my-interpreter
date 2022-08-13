package evaluator

import (
	"github.com/menxqk/my-interpreter/object"
)

func (e Evaluator) evalPlusOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s + %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Add(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s + %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalMinusOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s - %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Sub(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s - %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalAsteriskOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s * %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Mul(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s * %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalSlashOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s / %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Div(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s / %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalEqualOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s == %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Equ(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s == %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalNotEqualOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s != %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.NotEqu(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s != %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalGreaterOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s > %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Gt(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s > %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalGreaterEqualOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s >= %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Gte(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s >= %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalLesserOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s < %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Lt(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s < %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalLesserEqualOperator(left, right object.Object) object.Object {
	typeForObjects := getTypeForObjects(left, right)
	if typeForObjects == object.NullType {
		return newError("illegal operation %s <= %s", left.Type(), right.Type())
	}

	leftVal := left.ToType(typeForObjects)
	rightVal := right.ToType(typeForObjects)
	result := leftVal.Lte(rightVal)

	if result.Type() == object.NULL_OBJ {
		return newError("illegal operation %s <= %s", left.Type(), right.Type())
	}

	return result
}

func getTypeForObjects(left, right object.Object) object.ObjectType {

	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		return object.IntType
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		return object.FloatType
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ || left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		return object.FloatType
	}

	if left.Type() == object.CHAR_OBJ || left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ || right.Type() == object.STR_OBJ {
		return object.StringType
	}

	if left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ {
		return object.ArrayType
	}

	if left.Type() == object.BOOL_OBJ && right.Type() == object.BOOL_OBJ {
		return object.BooleanType
	}

	return object.NullType
}
