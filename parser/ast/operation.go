package ast

import (
	"fmt"
	"strings"

	lx "github.com/tamercuba/golisp/lexer"
)

type OperationNode struct {
	token  lx.Token
	Name   *Symbol
	Params []Node
}

func NewOperationNode(token lx.Token, Name *Symbol, Params []Node) *OperationNode {
	return &OperationNode{token, Name, Params}
}

func (o *OperationNode) GetToken() lx.Token {
	return o.token
}

func (o OperationNode) String() string {
	params := []string{}
	for _, param := range o.Params {
		params = append(params, fmt.Sprintf("%v", param))
	}
	return fmt.Sprintf("(%s %s)", o.Name.String(), strings.Join(params, " "))
}

func (o *OperationNode) GetValue() any {
	return o.Name
}

func IsValidOperation(name string) bool {
	switch name {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}
