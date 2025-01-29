package builtins

import (
	"fmt"

	"github.com/tamercuba/golisp/evaluator/object"
	"github.com/tamercuba/golisp/parser/ast"
)

func evalNumericListOp(o *ast.OperationNode, operation func(float64, float64) float64) (object.Object, error) {
	var resFloat float64
	hasFloat := false
	var err error = nil

	switch v := o.Params[0].(type) {
	case *ast.IntLiteral:
		resFloat = float64(v.GetValue().(int32))
	case *ast.FloatLiteral:
		resFloat = v.GetValue().(float64)
	default:
		return nil, NewBuiltinError(fmt.Sprintf("type isn't accepted in %s operation", o.Name), v.GetToken())

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
			err = NewBuiltinError(fmt.Sprintf("type isn't accepted in %s operation", o.Name), v.GetToken())
		}
	}

	if hasFloat {
		return &object.Float{Value: resFloat}, err
	} else {
		return &object.Integer{Value: int32(resFloat)}, err
	}
}

func EvalSum(o *ast.OperationNode) (object.Object, error) {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a + b
	})
}

func EvalSub(o *ast.OperationNode) (object.Object, error) {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a - b
	})
}

func EvalMultiplication(o *ast.OperationNode) (object.Object, error) {
	return evalNumericListOp(o, func(a, b float64) float64 {
		return a * b
	})
}
