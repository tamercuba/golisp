package evaluator

import (
	"fmt"

	"github.com/tamercuba/golisp/evaluator/builtins"
	"github.com/tamercuba/golisp/evaluator/object"
	"github.com/tamercuba/golisp/parser/ast"
)

type Evaluator struct {
	Env *Envinronment
}

func NewEvaluator() *Evaluator {
	return &Evaluator{Env: NewEnvironment()}
}

func (e *Evaluator) NewScope() {
	e.Env = e.Env.newScope()
}

func (e *Evaluator) DropScope() {
	e.Env = e.Env.dropScope()
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
	case *ast.OperationNode:
		return e.evalOperation(v)
	default:
		return nil
	}
}

func (e *Evaluator) evalOperation(o *ast.OperationNode) object.Object {
	e.evalOperationParams(o)
	switch o.Name.String() {
	case "+":
		if len(o.Params) < 2 {
			panic("+ needs 2 or more arguments")
		}
		return builtins.EvalSum(o)
	case "-":
		if len(o.Params) < 2 {
			panic("- needs 2 or more arguments")
		}
		return builtins.EvalSub(o)
	case "*":
		if len(o.Params) < 2 {
			panic("* needs 2 or more arguments")
		}
		return builtins.EvalMultiplication(o)
		// case "/":
	// panic("Not implemented yet")
	default:
		return &object.Nil{}
	}

}

func (e *Evaluator) evalOperationParams(o *ast.OperationNode) {
	newParams := []ast.Node{}
	for _, p := range o.Params {
		switch p.(type) {
		case *ast.Symbol:
			v := e.Env.Get(p.GetValue().(string))
			if v != nil {
				newParams = append(newParams, v)
			}
		default:
			newParams = append(newParams, p)
		}
	}

	o.Params = newParams
}

func (e *Evaluator) evalSymbol(l *ast.Symbol) object.Object {
	v := e.Env.Get(l.GetValue().(string))
	if v != nil {
		return e.evalNode(v)
	}
	panic(fmt.Sprintf("Unknow %s", v))
}

func (e *Evaluator) evalList(l *ast.ListExpression) object.Object {
	if l.Head == nil {
		return &object.List{Content: []object.Object{}}
	}

	switch s := l.Head.LNode.(type) {
	case *ast.Symbol:
		v := e.Env.Get(s.GetValue().(string))
		if v != nil {
			switch v.(type) {
			case *ast.LambdaNode:
				lambda := v.(*ast.LambdaNode)
				if ok, args := isLambdaCall(lambda, l.Head.Next); ok {
					return e.evalLambdaCall(lambda.Body, args, lambda.Args)
				}
				return &object.Nil{}
			}
		}
	case *ast.LambdaNode:
		currArgs := l.Head.Next
		params := []ast.Node{}

		for currArgs != nil {
			params = append(params, currArgs.LNode)
			currArgs = currArgs.Next
		}

		if len(params) != len(s.Args) {
			panic("WRONG NUMBER OF ARGUMENTS")
		}
		return e.evalLambdaCall(s.Body, params, s.Args)
	}

	c := l.Head
	r := []object.Object{}
	for c != nil {
		r = append(r, e.evalNode(c.LNode))
		c = c.Next
	}

	return &object.List{Content: r}
}

func isLambdaCall(l *ast.LambdaNode, n *ast.ListNode) (bool, []ast.Node) {
	arg := n
	totalArgs := 0
	args := []ast.Node{}
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
		err := e.Env.Bind(v.Name.GetValue().(string), v.Value)
		if err != nil {
			panic(err)
		}
	case ast.DEFINE:
		err := e.Env.BindGlobal(v.Name.GetValue().(string), v.Value)
		if err != nil {
			panic(err)
		}
	}
	return &object.Nil{}
}

func (e *Evaluator) evalLambdaCall(body ast.Node, params []ast.Node, args []ast.Symbol) object.Object {
	e.NewScope()

	if len(args) != len(params) {
		panic("AAA")
		// TODO: Improve validations
	}

	for i := range args {
		err := e.Env.Bind(args[i].GetValue().(string), params[i])
		if err != nil {
			panic(err)
		}
	}

	result := e.evalNode(body)
	e.DropScope()
	return result
}
