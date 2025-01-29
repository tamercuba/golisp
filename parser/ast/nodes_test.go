package ast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	lx "github.com/tamercuba/golisp/lexer"
)

func getToken(l string, tt lx.TokenType, ch int, col int) lx.Token {
	var t lx.Token
	t.Literal = l
	t.Type = tt
	t.Pos = *lx.NewPos(ch, col)
	return t
}

func TestDisplaySymbol(t *testing.T) {
	symbol := NewSymbol(getToken("a", lx.Symbol, 0, 0))

	assert.Equal(t, "a", fmt.Sprintf("%v", symbol))
	assert.Equal(t, "a", symbol.GetValue())
}

func TestStringNode(t *testing.T) {
	s := NewStringLiteral(getToken("\"a\"", lx.String, 0, 0))

	assert.Equal(t, "\"a\"", s.GetToken().Literal)
	assert.Equal(t, "'a'", fmt.Sprintf("%v", s))
	assert.Equal(t, "a", s.GetValue())
}

func TestIntegerLiteral(t *testing.T) {
	n := NewIntLiteral(getToken("1", lx.Int, 0, 0))

	assert.Equal(t, int32(1), n.GetValue())
	assert.Equal(t, "1", fmt.Sprintf("%v", n))
	assert.Equal(t, "1", n.GetToken().Literal)
}

func TestFloatLiteral(t *testing.T) {
	n := NewFloatLiteral(getToken("1.1", lx.Float, 0, 0))

	assert.Equal(t, float64(1.1), n.GetValue())
	assert.Equal(t, "1.100000f", fmt.Sprintf("%v", n))
	assert.Equal(t, "1.1", n.GetToken().Literal)
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

func TestBooleanDeclaration(t *testing.T) {
	b := NewBoolean(getToken("true", lx.Bool, 0, 0))

	assert.Equal(t, "true", fmt.Sprintf("%v", b))
	assert.Equal(t, true, b.GetValue())
	assert.Equal(t, "true", b.GetToken().Literal)
}

func TestNilNode(t *testing.T) {
	n := NewVoidNode(getToken("nil", lx.Void, 0, 0))

	assert.Equal(t, "nil", fmt.Sprintf("%v", n))
	assert.Nil(t, n.GetValue())
	assert.Equal(t, "nil", n.GetToken().Literal)
}

func TestLambdaNode(t *testing.T) {
	args := []Symbol{}

	body := NewListExpression(getToken("(", lx.LParen, 0, 0))
	l := NewLambdaNode(getToken("lambda", lx.Symbol, 0, 0), args, body)

	assert.Equal(t, "(lambda () ())", fmt.Sprintf("%v", l))
}

func TestVarDifinitionNode(t *testing.T) {
	name := NewSymbol(getToken("x", lx.Symbol, 0, 0))
	value := NewBoolean(getToken("true", lx.Bool, 0, 0))

	n := NewVarDifinitionNode(getToken("define", lx.Symbol, 0, 0), name, value)
	assert.Equal(t, "(define x true)", fmt.Sprintf("%v", n))
}

func TestOperationNode(t *testing.T) {
	name := NewSymbol(getToken("+", lx.Symbol, 0, 0))
	params := []Node{
		NewIntLiteral(getToken("1", lx.Int, 0, 0)),
		NewIntLiteral(getToken("2", lx.Int, 0, 0)),
	}
	on := NewOperationNode(getToken("(", lx.Symbol, 0, 0), name, params)

	assert.Equal(t, "(+ 1 2)", fmt.Sprintf("%v", on))
	assert.Equal(t, "+", fmt.Sprintf("%+v", on.GetValue()))
}
