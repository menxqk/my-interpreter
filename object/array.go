package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Array struct {
	ArrType  string
	Size     int
	Elements []Object
}

func (a *Array) Type() string { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s[%d] [", strings.ToLower(a.ArrType), a.Size))
	elems := []string{}
	for _, e := range a.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}
func (a *Array) ToType(objType ObjectType) Object {
	switch objType {
	case ArrayType:
		return a
	default:
		return &Null{}
	}
}
func (a *Array) Add(o Object) Object {
	oArr := o.(*Array)

	elems := a.Elements[:]
	for _, elem := range oArr.Elements {
		addElem := true
		for _, e := range a.Elements {
			if e.Inspect() == elem.Inspect() {
				addElem = false
				break
			}
		}
		if addElem {
			elems = append(elems, elem)
		}
	}

	return &Array{ArrType: a.ArrType, Elements: elems, Size: len(elems)}
}
func (a *Array) Sub(o Object) Object {
	oArr := o.(*Array)

	elems := []Object{}
	for _, elem := range a.Elements {
		addElem := true
		for _, e := range oArr.Elements {
			if e.Inspect() == elem.Inspect() {
				addElem = false
				break
			}
		}
		if addElem {
			elems = append(elems, elem)
		}
	}

	return &Array{ArrType: a.ArrType, Elements: elems, Size: len(elems)}
}
func (a *Array) Mul(o Object) Object { return &Null{} }
func (a *Array) Div(o Object) Object { return &Null{} }
func (a *Array) Equ(o Object) Object {
	return &Boolean{Value: a == o.(*Array)}
}
func (a *Array) NotEqu(o Object) Object {
	return &Boolean{Value: a != o.(*Array)}
}
func (a *Array) Gt(o Object) Object {
	return &Boolean{Value: a.Size > o.(*Array).Size}
}
func (a *Array) Gte(o Object) Object {
	return &Boolean{Value: a.Size >= o.(*Array).Size}
}
func (a *Array) Lt(o Object) Object {
	return &Boolean{Value: a.Size < o.(*Array).Size}
}
func (a *Array) Lte(o Object) Object {
	return &Boolean{Value: a.Size <= o.(*Array).Size}
}
