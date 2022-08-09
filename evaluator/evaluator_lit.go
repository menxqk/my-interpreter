package evaluator

import (
	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/object"
)

func (e *Evaluator) evalArrayLiteral(lit *ast.ArrayLiteral) object.Object {
	array := &object.Array{}
	array.Elements = []object.Object{}

	for _, elem := range lit.Elements {
		obj := e.Eval(elem)
		array.Elements = append(array.Elements, obj)
	}

	return array
}
