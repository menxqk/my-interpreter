package object

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() string                     { return RET_VAL_OBJ }
func (rv *ReturnValue) Inspect() string                  { return rv.Value.Inspect() }
func (rv *ReturnValue) ToType(objType ObjectType) Object { return &Null{} }
func (rv *ReturnValue) Add(o Object) Object              { return &Null{} }
func (rv *ReturnValue) Sub(o Object) Object              { return &Null{} }
func (rv *ReturnValue) Mul(o Object) Object              { return &Null{} }
func (rv *ReturnValue) Div(o Object) Object              { return &Null{} }
func (rv *ReturnValue) Equ(o Object) Object              { return &Null{} }
func (rv *ReturnValue) NotEqu(o Object) Object           { return &Null{} }
func (rv *ReturnValue) Gt(o Object) Object               { return &Null{} }
func (rv *ReturnValue) Gte(o Object) Object              { return &Null{} }
func (rv *ReturnValue) Lt(o Object) Object               { return &Null{} }
func (rv *ReturnValue) Lte(o Object) Object              { return &Null{} }
