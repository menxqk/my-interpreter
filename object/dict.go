package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Dict struct {
	Elements map[string]Object
}

func (d *Dict) Type() string { return DICT_OBJ }
func (d *Dict) Inspect() string {
	var out bytes.Buffer
	out.WriteString("dict{")
	elems := []string{}
	for k, v := range d.Elements {
		elems = append(elems, fmt.Sprintf("\"%s\": %s", k, v.Inspect()))
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("}")
	return out.String()
}
func (d *Dict) ToType(objType ObjectType) Object {
	switch objType {
	case DictType:
		return d
	default:
		return &Null{}
	}
}
func (d *Dict) Add(o Object) Object {
	oDict := o.(*Dict)

	elems := map[string]Object{}
	for k, v := range d.Elements {
		elems[k] = v
	}
	for k, v := range oDict.Elements {
		_, present := elems[k]
		if !present {
			elems[k] = v
		}
	}

	return &Dict{Elements: elems}
}
func (d *Dict) Sub(o Object) Object {
	oDict := o.(*Dict)

	elems := map[string]Object{}
	for k, v := range d.Elements {
		elems[k] = v
	}
	for k := range oDict.Elements {
		_, present := elems[k]
		if present {
			delete(elems, k)
		}
	}

	return &Dict{Elements: elems}
}
func (d *Dict) Mul(o Object) Object { return &Null{} }
func (d *Dict) Div(o Object) Object { return &Null{} }
func (d *Dict) Equ(o Object) Object {
	return &Boolean{Value: d == o.(*Dict)}
}
func (d *Dict) NotEqu(o Object) Object {
	return &Boolean{Value: d != o.(*Dict)}
}
func (d *Dict) Gt(o Object) Object {
	return &Boolean{Value: len(d.Elements) > len(o.(*Dict).Elements)}
}
func (d *Dict) Gte(o Object) Object {
	return &Boolean{Value: len(d.Elements) >= len(o.(*Dict).Elements)}
}
func (d *Dict) Lt(o Object) Object {
	return &Boolean{Value: len(d.Elements) < len(o.(*Dict).Elements)}
}
func (d *Dict) Lte(o Object) Object {
	return &Boolean{Value: len(d.Elements) <= len(o.(*Dict).Elements)}
}
