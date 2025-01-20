package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstToken(t *testing.T) {
	input := `(+ 1 2)`
	expected := Token{Type: LParen, Literal: "(", Pos: *newPos(0, 0)}

	lexer := NewLexer(input)
	firstToken := lexer.NextToken()

	assert.IsType(t, Token{}, firstToken)
	assert.Equal(t, expected, firstToken)
}

func TestDefunBasicDeclaration(t *testing.T) {
	input := `(defun x abc)`

	tests := []struct {
		expectedType   TokenType
		exeptedLiteral string
		posCh          int
		posCol         int
	}{
		{LParen, "(", 0, 0},
		{ReservedExpr, "defun", 1, 0},
		{Expr, "x", 7, 0},
		{Expr, "abc", 9, 0},
		{RParen, ")", 12, 0},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Pos.Ch, tt.posCh)
		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.exeptedLiteral)
	}
}

func TestPositionTracker(t *testing.T) {
	input := `(1 2 3)
(4 5 6)`
	tests := []struct {
		literal string
		ch      int
		col     int
	}{
		{"(", 0, 0},
		{"1", 1, 0},
		{"2", 3, 0},
		{"3", 5, 0},
		{")", 6, 0},
		{"(", 0, 1},
		{"4", 1, 1},
		{"5", 3, 1},
		{"6", 5, 1},
		{")", 6, 1},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Literal, tt.literal)
		assert.Equal(t, tok.Pos.Ch, tt.ch)
		assert.Equal(t, tok.Pos.Col, tt.col)
	}

}

func TestDefunDeclarationWithDashChar(t *testing.T) {
	input := `(defun variable-with-dash 10)`

	tests := []struct {
		expectedType   TokenType
		exeptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "defun"},
		{Expr, "variable-with-dash"},
		{Int, "10"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.exeptedLiteral)

	}
}

func TestFloatDeclaration(t *testing.T) {
	input := `(defun x 10.1)`

	tests := []struct {
		expectedType   TokenType
		exeptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "defun"},
		{Expr, "x"},
		{Float, "10.1"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.exeptedLiteral)

	}
}

func TestPlusExprAreRecognized(t *testing.T) {
	input := `(+ 1 2)`

	tests := []struct {
		expectedType   TokenType
		exeptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "+"},
		{Int, "1"},
		{Int, "2"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.exeptedLiteral)
	}
}
