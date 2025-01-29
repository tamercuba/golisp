package evaluator

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
)

type EvalError struct {
	Token lx.Token
	Msg   string
}

func NewEvalError(msg string, tok lx.Token) *EvalError {
	return &EvalError{Token: tok, Msg: msg}
}

func (e *EvalError) Error() string {
	return fmt.Sprintf("%v %v   %s", e.Token.Pos, e.Token.Literal, e.Msg)
}
