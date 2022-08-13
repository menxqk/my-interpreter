package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() string    { return INT_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) ToType(objType ObjectType) Object {
	switch objType {
	case IntType:
		return i
	case FloatType:
		return &Float{Value: float64(i.Value)}
	default:
		return &Null{}
	}
}
func (i *Integer) Add(o Object) Object {
	return &Integer{Value: i.Value + o.(*Integer).Value}
}
func (i *Integer) Sub(o Object) Object {
	return &Integer{Value: i.Value - o.(*Integer).Value}
}
func (i *Integer) Mul(o Object) Object {
	return &Integer{Value: i.Value * o.(*Integer).Value}
}
func (i *Integer) Div(o Object) Object {
	return &Integer{Value: i.Value / o.(*Integer).Value}
}
func (i *Integer) Equ(o Object) Object {
	return &Boolean{Value: i.Value == o.(*Integer).Value}
}
func (i *Integer) NotEqu(o Object) Object {
	return &Boolean{Value: i.Value != o.(*Integer).Value}
}
func (i *Integer) Gt(o Object) Object {
	return &Boolean{Value: i.Value > o.(*Integer).Value}
}
func (i *Integer) Gte(o Object) Object {
	return &Boolean{Value: i.Value >= o.(*Integer).Value}
}
func (i *Integer) Lt(o Object) Object {
	return &Boolean{Value: i.Value < o.(*Integer).Value}
}
func (i *Integer) Lte(o Object) Object {
	return &Boolean{Value: i.Value <= o.(*Integer).Value}
}
