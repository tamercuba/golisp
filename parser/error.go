package parser

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
)

type ParseError struct {
	Token lx.Token
	Msg   string
}

func NewParseError(Msg string, Token lx.Token) *ParseError {
	return &ParseError{Token, Msg}
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%v %v   %s", p.Token.Pos, p.Token.Literal, p.Msg)
}
