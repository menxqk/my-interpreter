package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/menxqk/my-interpreter/ast"
)

const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"

	INT_OBJ   = "INTEGER"
	FLOAT_OBJ = "FLOAT"
	CHAR_OBJ  = "CHAR"
	STR_OBJ   = "STRING"

	BOOL_OBJ    = "BOOLEAN"
	RET_VAL_OBJ = "RETURN_VALUE"
	FN_OBJ      = "FUNCTION"
)

type Object interface {
	Type() string
	Inspect() string
}

// NULL Object
type Null struct{}

func (n *Null) Type() string    { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

// ERROR Object
type Error struct {
	Message string
}

func (e *Error) Type() string    { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// INTEGER Object
type Integer struct {
	Value int64
}

func (i *Integer) Type() string    { return INT_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// FLOAT Object
type Float struct {
	Value float64
}

func (f *Float) Type() string    { return STR_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%.6f", f.Value) }

// CHAR Object
type Char struct {
	Value rune
}

func (c *Char) Type() string    { return CHAR_OBJ }
func (c *Char) Inspect() string { return fmt.Sprintf("%s", string(c.Value)) }

// STRING Object
type String struct {
	Value string
}

func (s *String) Type() string    { return STR_OBJ }
func (s *String) Inspect() string { return s.Value }

// BOOLEAN Object
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string    { return BOOL_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// RETURN VALUE Object
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() string    { return RET_VAL_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// FUNCTION Object
type Function struct {
	Identifier ast.Identifier
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() string { return FN_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s(", f.Identifier.Name))
	params := []string{}
	for _, param := range f.Parameters {
		params = append(params, param.String())
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}
