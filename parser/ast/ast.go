package ast

import (
	lx "github.com/tamercuba/golisp/lexer"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	ListStatements []ListExpression
}

type Identifier struct {
	Token lx.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
