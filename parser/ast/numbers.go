package ast

import (
	"fmt"
	"strconv"

	lx "github.com/tamercuba/golisp/lexer"
)

type IntLiteral struct {
	token lx.Token
	value int32
}

type FloatLiteral struct {
	token lx.Token
	value float64
}

func NewIntLiteral(token lx.Token) *IntLiteral {
	value, _ := strconv.Atoi(token.Literal)
	// Validated on lexer

	return &IntLiteral{token: token, value: int32(value)}
}

func (il *IntLiteral) GetToken() lx.Token {
	return il.token
}

func (il IntLiteral) String() string {
	return fmt.Sprintf("%d", il.value)
}

func (il *IntLiteral) GetValue() any {
	return il.value
}

func NewFloatLiteral(token lx.Token) *FloatLiteral {
	value, _ := strconv.ParseFloat(token.Literal, 64)
	// Validated on lexer

	return &FloatLiteral{token: token, value: value}
}

func (fl *FloatLiteral) GetToken() lx.Token {
	return fl.token
}

func (fl FloatLiteral) String() string {
	return fmt.Sprintf("%ff", fl.value)
}

func (fl *FloatLiteral) GetValue() any {
	return fl.value
}
