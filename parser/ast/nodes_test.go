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
	assert.Equal(t, "1", n.GetToken().Literal)
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
	assert.Equal(t, "1.1", n.GetToken().Literal)
}

func TestInvalidFloatLiteral(t *testing.T) {
	tok := lx.NewToken(0x61, lx.Int, 0, 0)
	assert.Panics(t, func() { NewFloatLiteral(tok) })
}

func TestListNode(t *testing.T) {
	tok := lx.NewToken(0x28, lx.Symbol, 0, 0)
	l := NewListExpression(tok)

	assert.Nil(t, l.Head)
	assert.Equal(t, "(", l.GetToken().Literal)

	i := NewSymbol(lx.NewToken(0x61, lx.Symbol, 0, 1))
	l.Append(i)

	assert.Equal(t, "(a)", fmt.Sprintf("%v", l))

	gv := l.GetValue()
	switch slice := gv.(type) {
	case []string:
		assert.Equal(t, 1, len(slice))
	default:
		assert.Fail(t, "Invalid type: %+v", slice)
	}
}

func TestFunctionDeclaration(t *testing.T) {
	tok := lx.NewToken(0x28, lx.Symbol, 0, 0)
	name := NewSymbol(tok)
	args := []Symbol{}
	body := Symbol{}
	l := NewFunctionDeclaration(tok, name, args, &body)

	assert.Equal(t, "(defun ( () ())", fmt.Sprintf("%v", l))
	assert.Equal(t, "(", l.GetValue())
	assert.Equal(t, "(", l.GetToken().Literal)
}

func TestBooleanDeclaration(t *testing.T) {
	var tok lx.Token
	tok.Literal = "true"
	tok.Type = lx.Bool
	b := NewBoolean(tok)

	assert.Equal(t, "true", fmt.Sprintf("%v", b))
	assert.Equal(t, true, b.GetValue())
	assert.Equal(t, "true", b.GetToken().Literal)
}

func TestNilNode(t *testing.T) {
	var tok lx.Token
	tok.Literal = "nil"
	tok.Type = lx.Void
	n := NewVoidNode(tok)

	assert.Equal(t, "nil", fmt.Sprintf("%v", n))
	assert.Nil(t, n.GetValue())
	assert.Equal(t, "nil", n.GetToken().Literal)
}
