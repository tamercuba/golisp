package ast

import (
	"strconv"

	lx "github.com/tamercuba/golisp/lexer"
)

type Boolean struct {
	token lx.Token
	value bool
}

func NewBoolean(token lx.Token) *Boolean {
	value, _ := strconv.ParseBool(token.Literal)
	// Error is unreachable because Lexer has already validated
	return &Boolean{token, value}
}

func (b *Boolean) GetToken() lx.Token {
	return b.token
}

func (b Boolean) String() string {
	return b.token.Literal
}

func (b *Boolean) GetValue() any {
	return b.value
}
