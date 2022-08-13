package object

type Error struct {
	Message string
}

func (e *Error) Type() string                     { return ERROR_OBJ }
func (e *Error) Inspect() string                  { return "ERROR: " + e.Message }
func (e *Error) ToType(objType ObjectType) Object { return nil }
func (e *Error) Add(o Object) Object              { return e }
func (e *Error) Sub(o Object) Object              { return e }
func (e *Error) Mul(o Object) Object              { return e }
func (e *Error) Div(o Object) Object              { return e }
