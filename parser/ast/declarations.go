package ast

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
)

type LetDeclaration struct {
	token lx.Token
	Name  *Symbol
	Value Node
}

func NewLetDeclaration(token lx.Token, name *Symbol, value Node) *LetDeclaration {
	return &LetDeclaration{token, name, value}
}

func (ld *LetDeclaration) GetToken() lx.Token {
	return ld.token
}

func (ld LetDeclaration) String() string {
	return fmt.Sprintf("(let %v %v)", ld.Name, ld.Value)
}

func (ld *LetDeclaration) GetValue() any {
	return ld.Value
}
