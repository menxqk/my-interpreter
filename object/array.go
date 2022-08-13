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
	out.WriteString(fmt.Sprintf("%s[%d] {", strings.ToLower(a.ArrType), a.Size))
	elems := []string{}
	for _, e := range a.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("}")
	return out.String()
}
func (a *Array) ToType(objType ObjectType) Object { return &Null{} }
func (a *Array) Add(o Object) Object              { return &Null{} }
func (a *Array) Sub(o Object) Object              { return &Null{} }
func (a *Array) Mul(o Object) Object              { return &Null{} }
func (a *Array) Div(o Object) Object              { return &Null{} }
