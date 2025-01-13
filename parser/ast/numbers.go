package ast

import (
	"fmt"
	lx "github.com/tamercuba/golisp/lexer"
)

type IntLiteral struct {
	Token lx.Token
	Value int32
}

type FloatLiteral struct {
	Token lx.Token
	Value float64
}

func (il *IntLiteral) expressionNode() {}

func (il *IntLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

func (fl *FloatLiteral) expressionNode() {}

func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fmt.Sprintf("%ff", fl.Value)
}
