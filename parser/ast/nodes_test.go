package ast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	lx "github.com/tamercuba/golisp/lexer"
)

func TestDisplaySymbol(t *testing.T) {
	tok := lx.NewToken(0x61, lx.Symbol, 0, 0)
	symbol := NewSymbol(tok)

	assert.Equal(t, "a", fmt.Sprintf("%v", symbol))
	assert.Equal(t, "a", symbol.GetValue())
}

func TestStringNode(t *testing.T) {
	var tok lx.Token
	tok.Literal = "\"a\""
	tok.Pos = *lx.NewPos(0, 0)
	s := NewStringLiteral(tok)

	assert.Equal(t, "\"a\"", s.GetToken().Literal)
	assert.Equal(t, "'a'", fmt.Sprintf("%v", s))
	assert.Equal(t, "a", s.GetValue())
}

func TestIntegerLiteral(t *testing.T) {
	tok := lx.NewToken(0x31, lx.Int, 0, 0)
	n := NewIntLiteral(tok)

	assert.Equal(t, int32(1), n.GetValue())
	assert.Equal(t, "1", fmt.Sprintf("%v", n))
}

func TestInvalidIntegerLiteral(t *testing.T) {
	tok := lx.NewToken(0x61, lx.Int, 0, 0)
	assert.Panics(t, func() { NewIntLiteral(tok) })
}

func TestFloatLiteral(t *testing.T) {
	var tok lx.Token
	tok.Literal = "1.1"
	tok.Pos = *lx.NewPos(0, 0)
	n := NewFloatLiteral(tok)

	assert.Equal(t, float64(1.1), n.GetValue())
	assert.Equal(t, "1.100000f", fmt.Sprintf("%v", n))
}

func TestInvalidFloatLiteral(t *testing.T) {
	tok := lx.NewToken(0x61, lx.Int, 0, 0)
	assert.Panics(t, func() { NewFloatLiteral(tok) })
}
