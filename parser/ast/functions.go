package ast

import (
	"fmt"
	"strings"

	lx "github.com/tamercuba/golisp/lexer"
)

type CallExpression struct {
	Token     lx.Token
	Function  Identifier // ?
	Arguments []Node
}

func (ce *CallExpression) expressionNode() {
	// This function should be empty for now
}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("(%s %s)", ce.Token.Literal, strings.Join(args, " "))
}

func (ce *CallExpression) GetValue() any {
	return ce.Function.GetValue()
}
