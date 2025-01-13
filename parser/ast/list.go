package ast

import (
	"fmt"
	lx "github.com/tamercuba/golisp/lexer"
	"strings"
)

type ListExpression struct {
	Token    lx.Token
	Elements []Expression
}

func (le *ListExpression) expressionNode() {}

func (le *ListExpression) TokenLiteral() string {
	return le.Token.Literal
}

func (le *ListExpression) String() string {
	values := []string{}
	for _, value := range le.Elements {
		values = append(values, value.String())
	}

	return fmt.Sprintf("(%s)", strings.Join(values, " "))
}
