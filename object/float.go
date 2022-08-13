package object

import "fmt"

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
func (f *Float) Add(o Object) Object { return &Float{Value: f.Value + o.(*Float).Value} }
func (f *Float) Sub(o Object) Object { return &Float{Value: f.Value - o.(*Float).Value} }
func (f *Float) Mul(o Object) Object { return &Float{Value: f.Value * o.(*Float).Value} }
func (f *Float) Div(o Object) Object { return &Float{Value: f.Value / o.(*Float).Value} }
