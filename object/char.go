package object

import "fmt"

type Char struct {
	Value rune
}

func (c *Char) Type() string    { return CHAR_OBJ }
func (c *Char) Inspect() string { return fmt.Sprintf("%s", string(c.Value)) }
func (c *Char) ToType(objType ObjectType) Object {
	switch objType {
	case CharType:
		return c
	case StringType:
		return &String{Value: string(c.Value)}
	default:
		return &Null{}
	}
}
func (c *Char) Add(o Object) Object {
	return &String{Value: string(c.Value) + o.(*String).Value}
}
func (c *Char) Sub(o Object) Object { return &Null{} }
func (c *Char) Mul(o Object) Object { return &Null{} }
func (c *Char) Div(o Object) Object { return &Null{} }
func (c *Char) Equ(o Object) Object {
	return &Boolean{Value: c.Value == o.(*Char).Value}
}
func (c *Char) NotEqu(o Object) Object {
	return &Boolean{Value: c.Value != o.(*Char).Value}
}
func (c *Char) Gt(o Object) Object {
	return &Boolean{Value: c.Value > o.(*Char).Value}
}
func (c *Char) Gte(o Object) Object {
	return &Boolean{Value: c.Value >= o.(*Char).Value}
}
func (c *Char) Lt(o Object) Object {
	return &Boolean{Value: c.Value < o.(*Char).Value}
}
func (c *Char) Lte(o Object) Object {
	return &Boolean{Value: c.Value <= o.(*Char).Value}
}
