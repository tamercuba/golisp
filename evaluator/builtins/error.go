package builtins

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
)

type BuiltinError struct {
	Token lx.Token
	Msg   string
}

func NewBuiltinError(Msg string, Token lx.Token) *BuiltinError {
	return &BuiltinError{Token, Msg}
}

func (b *BuiltinError) Error() string {
	return fmt.Sprintf("%v %v   %s", b.Token.Pos, b.Token.Literal, b.Msg)
}
