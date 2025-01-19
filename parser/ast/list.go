package ast

import (
	"fmt"
	lx "github.com/tamercuba/golisp/lexer"
	"strings"
)

type ListExpression struct {
	Token    lx.Token
	Elements []Node
}

func (le *ListExpression) expressionNode() {
	// This function should be empty for now
}

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

func (le *ListExpression) GetValue() any {
	result := make([]string, len(le.Elements))
	for _, el := range le.Elements {
		result = append(result, el.String())
	}

	return result
}
