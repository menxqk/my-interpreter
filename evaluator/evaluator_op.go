package evaluator

import (
	"math"

	"github.com/menxqk/my-interpreter/object"
)

func (e Evaluator) evalPlusOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Integer{Value: left.Value + right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Float{Value: left.Value + right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Float{Value: float64(left.Value) + right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Float{Value: left.Value + float64(right.Value)}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.String{Value: string(left.Value) + string(right.Value)}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.String{Value: left.Value + right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.String{Value: string(left.Value) + right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.String{Value: left.Value + string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s + %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalMinusOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Integer{Value: left.Value - right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Float{Value: left.Value - right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Float{Value: float64(left.Value) - right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Float{Value: left.Value - float64(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s - %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalAsteriskOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Integer{Value: left.Value * right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Float{Value: left.Value * right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Float{Value: float64(left.Value) * right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Float{Value: left.Value * float64(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s * %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalSlashOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Integer{Value: left.Value / right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Float{Value: left.Value / right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Float{Value: float64(left.Value) / right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Float{Value: left.Value / float64(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s / %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalEqualOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value == right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value == right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: isFloat64Equal(float64(left.Value), right.Value)}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: isFloat64Equal(left.Value, float64(right.Value))}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value == right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value == right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) == right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value == string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s == %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalNotEqualOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value != right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value != right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: !isFloat64Equal(float64(left.Value), right.Value)}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: !isFloat64Equal(left.Value, float64(right.Value))}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value != right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value != right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) != right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value != string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s != %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalGreaterOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value > right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value > right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: float64(left.Value) > right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value > float64(right.Value)}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value > right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value > right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) > right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value > string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s > %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalGreaterEqualOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value >= right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value >= right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: float64(left.Value) >= right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value >= float64(right.Value)}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value >= right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value >= right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) >= right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value >= string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s >= %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalLesserOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value < right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value < right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: float64(left.Value) < right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value < float64(right.Value)}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value < right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value < right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) < right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value < string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s < %s", left.Type(), right.Type())
	}

	return result
}

func (e Evaluator) evalLesserEqualOperator(left, right object.Object) object.Object {
	var result object.Object

	// Numbers
	if left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value <= right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Float)
		result = &object.Boolean{Value: left.Value <= right.Value}
	} else if left.Type() == object.INT_OBJ && right.Type() == object.FLOAT_OBJ {
		left := left.(*object.Integer)
		right := right.(*object.Float)
		result = &object.Boolean{Value: float64(left.Value) <= right.Value}
	} else if left.Type() == object.FLOAT_OBJ && right.Type() == object.INT_OBJ {
		left := left.(*object.Float)
		right := right.(*object.Integer)
		result = &object.Boolean{Value: left.Value <= float64(right.Value)}
	}

	// Chars and Strings
	if left.Type() == object.CHAR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value <= right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.String)
		right := right.(*object.String)
		result = &object.Boolean{Value: left.Value <= right.Value}
	} else if left.Type() == object.CHAR_OBJ && right.Type() == object.STR_OBJ {
		left := left.(*object.Char)
		right := right.(*object.String)
		result = &object.Boolean{Value: string(left.Value) <= right.Value}
	} else if left.Type() == object.STR_OBJ && right.Type() == object.CHAR_OBJ {
		left := left.(*object.String)
		right := right.(*object.Char)
		result = &object.Boolean{Value: left.Value <= string(right.Value)}
	}

	if result == nil {
		return newError("illegal operation %s <= %s", left.Type(), right.Type())
	}

	return result
}

const float64EqualityThreshold = 1e-10

func isFloat64Equal(left float64, right float64) bool {
	return math.Abs(left-right) <= float64EqualityThreshold
}
