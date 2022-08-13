package object

type ObjectType int

const (
	NullType = iota
	IntType
	FloatType
	CharType
	StringType
	ArrayType
	BoleanType
)

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"

	INT_OBJ   = "INT"
	FLOAT_OBJ = "FLOAT"
	CHAR_OBJ  = "CHAR"
	STR_OBJ   = "STRING"

	ARRAY_OBJ = "ARRAY"

	BOOL_OBJ    = "BOOLEAN"
	RET_VAL_OBJ = "RETURN_VALUE"
	FN_OBJ      = "FUNCTION"
)

type Object interface {
	Type() string
	Inspect() string
	ToType(ObjectType) Object
	Add(Object) Object
	Sub(Object) Object
	Mul(Object) Object
	Div(Object) Object

	// Equ(Object) Object
	// NotEqu(Object) Object
	// Gt(Object) Object
	// Gte(Object) Object
	// Lt(Object) Object
	// Lte(Object) Object
}

// Zero Value Object
func GetZeroValueObject(objType string) Object {
	switch objType {
	case INT_OBJ:
		return &Integer{}
	case FLOAT_OBJ:
		return &Float{}
	case CHAR_OBJ:
		return &Char{}
	case STR_OBJ:
		return &String{}
	case ARRAY_OBJ:
		return &Array{}
	default:
		return &Null{}
	}
}
