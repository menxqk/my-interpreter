package evaluator

const (
	NULL = "NULL_OBJ"
)

type Object interface {
	Type() string
	Value() string
}

type Null struct{}

func (n *Null) Type() string  { return "null" }
func (n *Null) Value() string { return NULL }
