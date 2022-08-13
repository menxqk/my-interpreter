package object

type Error struct {
	Message string
}

func (e *Error) Type() string                     { return ERROR_OBJ }
func (e *Error) Inspect() string                  { return "ERROR: " + e.Message }
func (e *Error) ToType(objType ObjectType) Object { return nil }
func (e *Error) Add(o Object) Object              { return &Null{} }
func (e *Error) Sub(o Object) Object              { return &Null{} }
func (e *Error) Mul(o Object) Object              { return &Null{} }
func (e *Error) Div(o Object) Object              { return &Null{} }
func (e *Error) Equ(o Object) Object {
	return &Boolean{Value: e == o.(*Error)}
}
func (e *Error) NotEqu(o Object) Object {
	return &Boolean{Value: e != o.(*Error)}
}
func (e *Error) Gt(o Object) Object  { return &Null{} }
func (e *Error) Gte(o Object) Object { return &Null{} }
func (e *Error) Lt(o Object) Object  { return &Null{} }
func (e *Error) Lte(o Object) Object { return &Null{} }
