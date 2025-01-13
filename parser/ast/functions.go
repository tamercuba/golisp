package ast

import (
	"fmt"
	"strings"

	lx "github.com/tamercuba/golisp/lexer"
)

type CallExpression struct {
	Token     lx.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("(%s %s)", ce.Function.String(), strings.Join(args, " "))
}
