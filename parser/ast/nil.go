package ast

import (
	lx "github.com/tamercuba/golisp/lexer"
)

type VoidNode struct {
	token lx.Token
}

func NewVoidNode(token lx.Token) *VoidNode {
	return &VoidNode{token}
}

func (v *VoidNode) GetToken() lx.Token {
	return v.token
}

func (v VoidNode) String() string {
	return "nil"
}

func (v *VoidNode) GetValue() any {
	return nil
}
