package ast

import (
	"fmt"
	"strings"

	lx "github.com/tamercuba/golisp/lexer"
)

type LambdaNode struct {
	token lx.Token
	Args  []Symbol
	Body  Node
}

func NewLambdaNode(token lx.Token, Args []Symbol, Body Node) *LambdaNode {
	return &LambdaNode{token, Args, Body}
}

func (l *LambdaNode) GetToken() lx.Token {
	return l.token
}

func (l LambdaNode) String() string {
	argsStrings := []string{}
	for _, argName := range l.Args {
		argsStrings = append(argsStrings, argName.String())
	}
	argsResult := fmt.Sprintf("%v", strings.Join(argsStrings, " "))
	return fmt.Sprintf("(lambda (%v) (%v))", argsResult, l.Body.String())
}

func (l *LambdaNode) GetValue() any {
	return nil
}
