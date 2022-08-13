package object

import (
	"fmt"
	"math"
)

type Float struct {
	Value float64
}

func (f *Float) Type() string    { return FLOAT_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%.6f", f.Value) }
func (f *Float) ToType(objType ObjectType) Object {
	switch objType {
	case IntType:
		return &Integer{Value: int64(f.Value)}
	case FloatType:
		return f
	default:
		return &Null{}
	}
}
func (f *Float) Add(o Object) Object {
	return &Float{Value: f.Value + o.(*Float).Value}
}
func (f *Float) Sub(o Object) Object {
	return &Float{Value: f.Value - o.(*Float).Value}
}
func (f *Float) Mul(o Object) Object {
	return &Float{Value: f.Value * o.(*Float).Value}
}
func (f *Float) Div(o Object) Object {
	return &Float{Value: f.Value / o.(*Float).Value}
}
func (f *Float) Equ(o Object) Object {
	return &Boolean{Value: isFloat64Equal(f.Value, o.(*Float).Value)}
}
func (f *Float) NotEqu(o Object) Object {
	return &Boolean{Value: !isFloat64Equal(f.Value, o.(*Float).Value)}
}
func (f *Float) Gt(o Object) Object {
	return &Boolean{Value: f.Value > o.(*Float).Value}
}
func (f *Float) Gte(o Object) Object {
	return &Boolean{Value: f.Value >= o.(*Float).Value}
}
func (f *Float) Lt(o Object) Object {
	return &Boolean{Value: f.Value < o.(*Float).Value}
}
func (f *Float) Lte(o Object) Object {
	return &Boolean{Value: f.Value <= o.(*Float).Value}
}

const float64EqualityThreshold = 1e-10

func isFloat64Equal(left float64, right float64) bool {
	return math.Abs(left-right) <= float64EqualityThreshold
}
