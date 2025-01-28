package builtins

import (
	"fmt"

	"github.com/tamercuba/golisp/evaluator/object"
	"github.com/tamercuba/golisp/parser/ast"
)

func evalNumericListOp(o *ast.OperationNode, operation func(float64, float64) float64) object.Object {
	var resFloat float64
	hasFloat := false

	switch v := o.Params[0].(type) {
	case *ast.IntLiteral:
		resFloat = float64(v.GetValue().(int32))
	case *ast.FloatLiteral:
		resFloat = v.GetValue().(float64)

	}

	for _, node := range o.Params[1:] {
		var valueFloat float64
		switch v := node.(type) {
		case *ast.IntLiteral:
			resFloat = operation(resFloat, float64(v.GetValue().(int32)))
		case *ast.FloatLiteral:
			valueFloat = float64(v.GetValue().(float64))
			hasFloat = true
			resFloat = operation(resFloat, valueFloat)
		case *ast.Symbol:
		default:
			panic(fmt.Sprintf("Unsupported type: %T", v))
		}
	}

	if hasFloat {
		return &object.Float{Value: resFloat}
	} else {
		return &object.Integer{Value: int32(resFloat)}
	}
}

func EvalSum(o *ast.OperationNode) object.Object {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a + b
	})
}

func EvalSub(o *ast.OperationNode) object.Object {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a - b
	})
}

func EvalMultiplication(o *ast.OperationNode) object.Object {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a * b
	})
}
