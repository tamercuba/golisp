package ast

import (
	"fmt"
	lx "github.com/tamercuba/golisp/lexer"
)

type StringLiteral struct {
	token lx.Token
	value string
}

func NewStringLiteral(token lx.Token) *StringLiteral {
	value := token.Literal[1 : len(token.Literal)-1]
	return &StringLiteral{token: token, value: value}
}

func (sl *StringLiteral) GetToken() lx.Token {
	return sl.token
}

func (sl StringLiteral) String() string {
	return fmt.Sprintf("'%s'", sl.value)
}

func (sl *StringLiteral) GetValue() any {
	return sl.value
}
