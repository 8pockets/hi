package evaluator

import (
	"github.com/8pockets/hi/ast"
	"github.com/8pockets/hi/object"
)

func evalBlockStatement(
	block *ast.BlockStatement,
	env *object.Environment,
) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

func evalAssignStatement(
	ae *ast.AssignStatement,
	env *object.Environment,
) object.Object {

	leftVal := ae.Left.String()
	left, ok := env.Get(leftVal)
	if !ok {
		return newError("undefined variable: %s", leftVal)
	}

	if left.IsAssign != true {
		return newError("Cannot assign %s", leftVal)
	}

	right := Eval(ae.Right, env)
	if isError(right) {
		return right
	}
	// type check
	if left.Obj.Type() != right.Type() {
		return newError("Cannot assign diferent type: %s and %s", left.Obj.Type(), right.Type())
	}

	env.Set(leftVal, object.ObjData{Obj: right, IsAssign: true})

	return NULL
}
