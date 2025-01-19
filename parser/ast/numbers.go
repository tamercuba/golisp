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

func (il *IntLiteral) expressionNode() {
	// This function should be empty for now
}

func (il *IntLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

func (il *IntLiteral) GetValue() any {
	return il.Value
}

func (fl *FloatLiteral) expressionNode() {
	// This function should be empty for now
}

func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fmt.Sprintf("%ff", fl.Value)
}

func (fl *FloatLiteral) GetValue() any {
	return fl.Value
}
