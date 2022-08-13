package object

type String struct {
	Value string
}

func (s *String) Type() string    { return STR_OBJ }
func (s *String) Inspect() string { return s.Value }
func (s *String) ToType(objType ObjectType) Object {
	switch objType {
	case CharType:
		if len(s.Value) > 0 {
			return &Char{Value: rune(s.Value[0])}
		}
		return &Char{Value: 0}
	case StringType:
		return s
	default:
		return &Null{}
	}
}
func (s *String) Add(o Object) Object {
	return &String{Value: s.Value + o.(*String).Value}
}
func (s *String) Sub(o Object) Object { return &Null{} }
func (s *String) Mul(o Object) Object { return &Null{} }
func (s *String) Div(o Object) Object { return &Null{} }
