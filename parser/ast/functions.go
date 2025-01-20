package ast

import (
	"fmt"
	"strings"

	lx "github.com/tamercuba/golisp/lexer"
)

type FunctionDeclaration struct {
	token lx.Token
	Name  *Symbol
	Args  []Symbol
	Body  Node
}

func NewFunctionDeclaration(token lx.Token, name *Symbol, args []Symbol, body Node) *FunctionDeclaration {
	return &FunctionDeclaration{token: token, Name: name, Args: args, Body: body}
}

func (fd *FunctionDeclaration) GetToken() lx.Token {
	return fd.token
}

func (fd FunctionDeclaration) String() string {
	argsStrings := []string{}
	for _, argName := range fd.Args {
		argsStrings = append(argsStrings, argName.String())
	}
	argsResult := fmt.Sprintf("%q", strings.Join(argsStrings, " "))
	return fmt.Sprintf("(defun %q (%q) (%q))", fd.Name, argsResult, fd.Body.String())
}

func (fd *FunctionDeclaration) GetValue() any {
	return fd.Name
}
