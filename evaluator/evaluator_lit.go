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
		if array.ArrType == "" {
			array.ArrType = obj.Type()
		}
		if obj.Type() != array.ArrType {
			return newError("cannot mix %s and %s in array", array.ArrType, obj.Type())
		}
		array.Elements = append(array.Elements, obj)
	}
	array.Size = len(array.Elements)

	return array
}
