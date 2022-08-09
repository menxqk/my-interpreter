package evaluator

import (
	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/object"
)

func (e *Evaluator) evalExpressionStatement(stmt *ast.ExpressionStatement) object.Object {
	return e.Eval(stmt.Expression)
}

func (e *Evaluator) evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
		result = e.Eval(stmt)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func (e *Evaluator) evalFunctionDeclarationStatement(stmt *ast.FunctionDeclarationStatement) object.Object {
	var result object.Object

	obj := e.Eval(stmt.Function)
	fn, ok := obj.(*object.Function)
	if !ok {
		return newError("could not eval function")
	}

	result = e.env.Set(fn.Identifier.Name, fn)

	return result
}

func (e *Evaluator) evalArrayDeclarationStatement(stmt *ast.ArrayDeclarationStatement) object.Object {
	var result object.Object

	name := stmt.Identifier.Name
	arrType := stmt.Identifier.Type
	size := stmt.Size

	obj := e.Eval(stmt.Expression)
	// if expression is null, set zero value Object for the type
	if obj.Type() == object.NULL_OBJ {
		obj = object.GetZeroValueObject(object.ARRAY_OBJ)
	}

	if obj.Type() != object.ARRAY_OBJ {
		return newError("cannot assign %s to %s", obj.Type(), object.ARRAY_OBJ)
	}

	arrObj := obj.(*object.Array)
	arrObj.ArrType = arrType
	arrObj.Size = size

	for _, elem := range arrObj.Elements {
		if elem.Type() != arrType {
			return newError("cannot assign %s to %s array", elem.Type(), arrType)
		}
	}

	if len(arrObj.Elements) > arrObj.Size {
		return newError("%d elements exceed array capacity %d", len(arrObj.Elements), arrObj.Size)
	}

	result = e.env.Set(name, obj)

	return result
}

func (e *Evaluator) evalVariableDeclarationStatement(stmt *ast.VariableDeclarationStatement) object.Object {
	var result object.Object

	name := stmt.Identifier.Name
	varType := stmt.Identifier.Type

	obj := e.Eval(stmt.Expression)
	// if expression is null, set zero value Object for the type
	if obj.Type() == object.NULL_OBJ {
		obj = object.GetZeroValueObject(varType)
	}

	if obj.Type() != varType {
		return newError("cannot assign %s to %s", obj.Type(), varType)
	}

	result = e.env.Set(name, obj)

	return result
}

func (e *Evaluator) evalAssignmentStatement(stmt *ast.AssignmentStatement) object.Object {
	var result object.Object

	obj, ok := e.env.Get(stmt.Identifier.Name)
	if !ok {
		return newError("%q not declared", stmt.Identifier.Name)
	}

	expObj := e.Eval(stmt.Expression)

	if expObj.Type() != obj.Type() {
		return newError("cannot assign %s to %s", expObj.Type(), obj.Type())
	}

	// if obj is array, check assignment array
	arrObj, isArray := obj.(*object.Array)
	if isArray {
		expObjArray := expObj.(*object.Array)

		// check element types of assignment array
		for _, elem := range expObjArray.Elements {
			if elem.Type() != arrObj.ArrType {
				return newError("cannot assign %s to %s array", elem.Type(), arrObj.ArrType)
			}
		}

		// check size of assignment array
		if len(expObjArray.Elements) > arrObj.Size {
			return newError("%d elements exceed array capacity %d", len(expObjArray.Elements), arrObj.Size)
		}

		// if checks are ok, set correct type and size for assignment array
		expObjArray.ArrType = arrObj.ArrType
		expObjArray.Size = arrObj.Size
	}

	result = e.env.Set(stmt.Identifier.Name, expObj)

	return result
}

func (e *Evaluator) evalReturnStatement(stmt *ast.ReturnStatement) object.Object {
	return e.Eval(stmt.ReturnValue)
}
