package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/menxqk/my-interpreter/ast"
)

type Function struct {
	Identifier ast.Identifier
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() string { return FN_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s %s(", f.Identifier.TypeLiteral, f.Identifier.Name))
	params := []string{}
	for _, param := range f.Parameters {
		params = append(params, param.String())
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}
func (f *Function) ToType(objType ObjectType) Object { return &Null{} }
func (f *Function) Add(o Object) Object              { return &Null{} }
func (f *Function) Sub(o Object) Object              { return &Null{} }
func (f *Function) Mul(o Object) Object              { return &Null{} }
func (f *Function) Div(o Object) Object              { return &Null{} }
func (f *Function) Equ(o Object) Object {
	return &Boolean{Value: f == o.(*Function)}
}
func (f *Function) NotEqu(o Object) Object {
	return &Boolean{Value: f != o.(*Function)}
}
func (f *Function) Gt(o Object) Object  { return &Null{} }
func (f *Function) Gte(o Object) Object { return &Null{} }
func (f *Function) Lt(o Object) Object  { return &Null{} }
func (f *Function) Lte(o Object) Object { return &Null{} }
