package ast

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
)

type DefType uint8

const (
	LET DefType = iota
	DEFINE
)

type VarDifinitionNode struct {
	token          lx.Token
	Name           *Symbol
	Value          Node
	DefinitionType DefType
}

func NewVarDifinitionNode(token lx.Token, name *Symbol, value Node) *VarDifinitionNode {
	var DefinitionType DefType
	if token.Literal == "let" {
		DefinitionType = LET
	} else {
		DefinitionType = DEFINE
	}
	return &VarDifinitionNode{token, name, value, DefinitionType}
}

func (ld *VarDifinitionNode) GetToken() lx.Token {
	return ld.token
}

func (ld VarDifinitionNode) String() string {
	return fmt.Sprintf("(%v %v %v)", ld.GetToken().Literal, ld.Name, ld.Value)
}

func (ld *VarDifinitionNode) GetValue() any {
	return ld.Value
}
