package evaluator

import (
	"fmt"

	"github.com/tamercuba/golisp/evaluator/builtins"
	"github.com/tamercuba/golisp/evaluator/object"
	lx "github.com/tamercuba/golisp/lexer"
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

func (e *Evaluator) EvalProgram(p *ast.Program) (object.Object, error) {
	var result object.Object
	var err error

	for _, node := range p.ListStatements {
		result, err = e.evalNode(node)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *Evaluator) evalNode(p ast.Node) (object.Object, error) {
	switch v := p.(type) {
	case *ast.IntLiteral:
		return &object.Integer{Value: v.GetValue().(int32)}, nil
	case *ast.FloatLiteral:
		return &object.Float{Value: v.GetValue().(float64)}, nil
	case *ast.Boolean:
		return &object.Boolean{Value: v.GetValue().(bool)}, nil
	case *ast.VoidNode:
		return &object.Nil{}, nil
	case *ast.StringLiteral:
		return &object.String{Value: v.GetValue().(string)}, nil
	case *ast.Symbol:
		return e.evalSymbol(v)
	case *ast.ListExpression:
		return e.evalList(v)
	case *ast.VarDifinitionNode:
		return e.evalVarDefinition(v)
	case *ast.LambdaNode:
		return &object.Nil{}, nil
	case *ast.OperationNode:
		return e.evalOperation(v)
	default:
		fmt.Println("NAO SEI")
		// TODO: Dont know what to do yet
		return nil, nil
	}
}

func (e *Evaluator) evalOperation(o *ast.OperationNode) (object.Object, error) {
	e.evalOperationParams(o)
	switch o.Name.String() {
	case "+":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalSum(o)
	case "-":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalSub(o)
	case "*":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalMultiplication(o)
	case "=":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalEqual(o)
	case "<":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalLesser(o)
	case ">":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalGreather(o)
	case ">=":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalGreatherOrEqual(o)
	case "<=":
		if len(o.Params) < 2 {
			return nil, NewEvalError(fmt.Sprintf("needs 2 or more arguments, %d were given", len(o.Params)), o.GetToken())
		}
		return builtins.EvalLesserOrEqual(o)
		// case "/":
	// panic("Not implemented yet")
	default:
		return nil, NewEvalError(fmt.Sprintf("Unknown symbol %s", o.Name.String()), o.GetToken())
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

func (e *Evaluator) evalSymbol(l *ast.Symbol) (object.Object, error) {
	v := e.Env.Get(l.GetValue().(string))
	if v != nil {
		return e.evalNode(v)
	}

	return nil, NewEvalError(fmt.Sprintf("Unknown symbol"), l.GetToken())
}

func (e *Evaluator) evalList(l *ast.ListExpression) (object.Object, error) {
	if l.Head == nil {
		return &object.List{Content: []object.Object{}}, nil
	}

	switch s := l.Head.LNode.(type) {
	case *ast.Symbol:
		v := e.Env.Get(s.GetValue().(string))
		if v != nil {
			switch v.(type) {
			case *ast.LambdaNode:
				lambda := v.(*ast.LambdaNode)
				if ok, args := isLambdaCall(lambda, l.Head.Next); ok {
					return e.evalLambdaCall(lambda.GetToken(), lambda.Body, args, lambda.Args)
				}
				return &object.Nil{}, nil
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
			return nil, NewEvalError(fmt.Sprintf(
				"%d arguments expected, %d given", len(s.Args), len(params),
			), l.GetToken())
		}
		return e.evalLambdaCall(s.GetToken(), s.Body, params, s.Args)
	}

	c := l.Head
	r := []object.Object{}
	for c != nil {
		n, err := e.evalNode(c.LNode)
		if err != nil {
			return nil, err
		}
		r = append(r, n)
		c = c.Next
	}

	return &object.List{Content: r}, nil
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

func (e *Evaluator) evalVarDefinition(v *ast.VarDifinitionNode) (object.Object, error) {
	switch v.DefinitionType {
	case ast.LET:
		err := e.Env.Bind(v.Name.GetValue().(string), v.Value)
		if err != nil {
			return nil, err
		}
	case ast.DEFINE:
		err := e.Env.BindGlobal(v.Name.GetValue().(string), v.Value)
		if err != nil {
			return nil, err
		}
	}
	return &object.Nil{}, nil
}

func (e *Evaluator) evalLambdaCall(t lx.Token, body ast.Node, params []ast.Node, args []ast.Symbol) (object.Object, error) {
	e.NewScope()

	if len(args) != len(params) {
		return nil, NewEvalError(fmt.Sprintf(
			"%d arguments expected, %d given", len(args), len(params),
		), t)

	}

	for i := range args {
		err := e.Env.Bind(args[i].GetValue().(string), params[i])
		if err != nil {
			return nil, err
		}
	}

	result, err := e.evalNode(body)
	e.DropScope()
	return result, err
}
