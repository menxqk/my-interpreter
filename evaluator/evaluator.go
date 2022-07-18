package evaluator

import (
	"github.com/menxqk/my-interpreter/ast"
)

type Evaluator struct {
}

func New() *Evaluator {
	e := &Evaluator{}
	return e
}

func (e *Evaluator) Eval(prog *ast.Program) Object {
	return &Null{}
}
