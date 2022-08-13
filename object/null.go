package object

type Null struct{}

func (n *Null) Type() string                     { return NULL_OBJ }
func (n *Null) Inspect() string                  { return "null" }
func (n *Null) ToType(objType ObjectType) Object { return n }
func (n *Null) Add(o Object) Object              { return n }
func (n *Null) Sub(o Object) Object              { return n }
func (n *Null) Mul(o Object) Object              { return n }
func (n *Null) Div(o Object) Object              { return n }
func (n *Null) Equ(o Object) Object              { return n }
func (n *Null) NotEqu(o Object) Object           { return n }
func (n *Null) Gt(o Object) Object               { return n }
func (n *Null) Gte(o Object) Object              { return n }
func (n *Null) Lt(o Object) Object               { return n }
func (n *Null) Lte(o Object) Object              { return n }
