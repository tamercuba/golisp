package ast

import (
	lx "github.com/tamercuba/golisp/lexer"
)

type Node interface {
	GetToken() lx.Token
	String() string
	GetValue() any
}

type Program struct {
	ListStatements []Node
}

type Symbol struct {
	token lx.Token
	value string
}

func NewSymbol(token lx.Token) *Symbol {
	return &Symbol{token: token, value: token.Literal}
}

func (s *Symbol) GetToken() lx.Token {
	return s.token
}

func (s Symbol) String() string {
	return s.value
}

func (s *Symbol) GetValue() any {
	return s.value
}
