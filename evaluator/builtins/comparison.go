package builtins

import (
	"github.com/tamercuba/golisp/evaluator/object"
	"github.com/tamercuba/golisp/parser/ast"
)

func compareList(o *ast.OperationNode, ff func(ast.Node, ast.Node) bool) object.Object {
	value := o.Params[0]

	for _, v := range o.Params[1:] {
		if !ff(value, v) {
			return &object.Boolean{Value: false}
		}
		value = v
	}

	return &object.Boolean{Value: true}
}

func EvalEqual(o *ast.OperationNode) object.Object {
	return compareList(o, func(a, b ast.Node) bool {
		return a.GetValue() == b.GetValue()
	})
}

func EvalLesser(o *ast.OperationNode) object.Object {
	return compareList(o, func(a, b ast.Node) bool {
		switch va := a.GetValue().(type) {
		case int32:
			if vb, ok := b.GetValue().(int32); ok {
				return va < vb
			}
		case float64:
			if vb, ok := b.GetValue().(float64); ok {
				return va < vb
			}
		}
		return false
	})
}

func EvalGreather(o *ast.OperationNode) object.Object {
	return compareList(o, func(a, b ast.Node) bool {
		switch va := a.GetValue().(type) {
		case int32:
			if vb, ok := b.GetValue().(int32); ok {
				return va > vb
			}
		case float64:
			if vb, ok := b.GetValue().(float64); ok {
				return va > vb
			}
		}
		return false
	})
}

func EvalGreatherOrEqual(o *ast.OperationNode) object.Object {
	return compareList(o, func(a, b ast.Node) bool {
		switch va := a.GetValue().(type) {
		case int32:
			if vb, ok := b.GetValue().(int32); ok {
				return va >= vb
			}
		case float64:
			if vb, ok := b.GetValue().(float64); ok {
				return va >= vb
			}
		}
		return false
	})
}

func EvalLesserOrEqual(o *ast.OperationNode) object.Object {
	return compareList(o, func(a, b ast.Node) bool {
		switch va := a.GetValue().(type) {
		case int32:
			if vb, ok := b.GetValue().(int32); ok {
				return va <= vb
			}
		case float64:
			if vb, ok := b.GetValue().(float64); ok {
				return va <= vb
			}
		}
		return false
	})
}
