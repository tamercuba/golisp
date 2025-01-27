package evaluator

import (
	"fmt"

	"github.com/tamercuba/golisp/evaluator/object"
	"github.com/tamercuba/golisp/parser/ast"
)

type Evaluator struct {
	env *Envinronment
}

func NewEvaluator() *Evaluator {
	return &Evaluator{env: NewEnvironment()}
}

func (e *Evaluator) EvalProgram(p *ast.Program) object.Object {
	var result object.Object

	for _, node := range p.ListStatements {
		result = e.evalNode(node)
	}

	return result
}

func (e *Evaluator) evalNode(p ast.Node) object.Object {
	switch v := p.(type) {
	case *ast.IntLiteral:
		return &object.Integer{Value: v.GetValue().(int32)}
	case *ast.FloatLiteral:
		return &object.Float{Value: v.GetValue().(float64)}
	case *ast.Boolean:
		return &object.Boolean{Value: v.GetValue().(bool)}
	case *ast.VoidNode:
		return &object.Nil{}
	case *ast.StringLiteral:
		return &object.String{Value: v.GetValue().(string)}
	case *ast.Symbol:
		return e.evalSymbol(v)
	case *ast.ListExpression:
		return e.evalList(v)
	case *ast.VarDifinitionNode:
		return e.evalVarDefinition(v)
	case *ast.LambdaNode:
		return &object.Nil{}
	default:
		return nil
	}
}

func (e *Evaluator) evalSymbol(l *ast.Symbol) object.Object {
	v := e.env.Get(l.GetValue().(string))
	if v != nil {
		return e.evalNode(v)
	}
	panic(fmt.Sprintf("Unknow %s", v))
}

func (e *Evaluator) evalList(l *ast.ListExpression) object.Object {
	if l.Head == nil {
		return &object.List{Content: []object.Object{}}
	}

	switch s := l.Head.Next.LNode.(type) {
	case *ast.Symbol:
		v := e.env.Get(s.GetValue().(string))
		if v != nil {
			switch v.(type) {
			case *ast.LambdaNode:
				lambda := v.(*ast.LambdaNode)
				if ok, args := isLambdaCall(lambda, l.Head.Next.Next); ok {
					return e.evalLambdaCall(lambda.Body, args, lambda.Args)
				}
				return &object.Nil{}
			}
		}
	}

	c := l.Head
	r := make([]object.Object, l.Size)
	for c != nil {
		r = append(r, e.evalNode(c.LNode))
		c = c.Next
	}
	return nil

}

func isLambdaCall(l *ast.LambdaNode, n *ast.ListNode) (bool, []ast.Node) {
	arg := n
	totalArgs := 0
	args := make([]ast.Node, len(l.Args))
	for arg != nil {
		totalArgs++
		args = append(args, arg.LNode)
		arg = arg.Next
	}

	return totalArgs == len(l.Args), args
}

func (e *Evaluator) evalVarDefinition(v *ast.VarDifinitionNode) object.Object {
	switch v.DefinitionType {
	case ast.LET:
		err := e.env.Bind(v.Name.GetValue().(string), v.Value)
		if err != nil {
			panic(err)
		}
	case ast.DEFINE:
		err := e.env.BindGlobal(v.Name.GetValue().(string), v.Value)
		if err != nil {
			panic(err)
		}
	}
	return &object.Nil{}
}

func (e *Evaluator) evalLambdaCall(body ast.Node, params []ast.Node, args []ast.Symbol) object.Object {
	e.env.NewScope()
	if len(args) != len(params) {
		panic("AAA")
		// TODO: Improve validations
	}

	for i := range args {
		err := e.env.Bind(args[i].GetValue().(string), params[i])
		if err != nil {
			panic(err)
		}
	}

	defer e.env.DropScope()
	return e.evalNode(body)
}
