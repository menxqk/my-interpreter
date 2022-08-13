package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string                     { return BOOL_OBJ }
func (b *Boolean) Inspect() string                  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) ToType(objType ObjectType) Object { return b }
func (b *Boolean) Add(o Object) Object              { return &Null{} }
func (b *Boolean) Sub(o Object) Object              { return &Null{} }
func (b *Boolean) Mul(o Object) Object              { return &Null{} }
func (b *Boolean) Div(o Object) Object              { return &Null{} }
func (b *Boolean) Equ(o Object) Object {
	return &Boolean{Value: b.Value == o.(*Boolean).Value}
}
func (b *Boolean) NotEqu(o Object) Object {
	return &Boolean{Value: b.Value != o.(*Boolean).Value}
}
func (b *Boolean) Gt(o Object) Object  { return &Null{} }
func (b *Boolean) Gte(o Object) Object { return &Null{} }
func (b *Boolean) Lt(o Object) Object  { return &Null{} }
func (b *Boolean) Lte(o Object) Object { return &Null{} }
