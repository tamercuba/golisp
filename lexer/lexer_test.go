package lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstToken(t *testing.T) {
	input := `(+ 1 2)`
	expected := Token{Type: LParen, Literal: "(", Pos: *NewPos(0, 0)}

	lexer := NewLexer(input)
	firstToken := lexer.NextToken()

	assert.IsType(t, Token{}, firstToken)
	assert.Equal(t, expected, firstToken)
}

func TestEOF(t *testing.T) {
	input := "\x00"
	expected := Token{Type: EOF, Literal: "\x00", Pos: *NewPos(0, 0)}

	l := NewLexer(input)
	tt := l.NextToken()

	assert.IsType(t, Token{}, tt)
	assert.Equal(t, expected, tt)
}

func TestIllegalToken(t *testing.T) {
	input := ","
	expected := Token{Type: IllegalToken, Literal: ",", Pos: *NewPos(0, 0)}

	l := NewLexer(input)
	tt := l.NextToken()

	assert.IsType(t, Token{}, tt)
	assert.Equal(t, expected, tt)

}

func TestDefunBasicDeclaration(t *testing.T) {
	input := `(defun x abc)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
		posCh           int
		posCol          int
	}{
		{LParen, "(", 0, 0},
		{Symbol, "defun", 1, 0},
		{Symbol, "x", 7, 0},
		{Symbol, "abc", 9, 0},
		{RParen, ")", 12, 0},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Pos.Ch, tt.posCh)
		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)
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
		expectedType    TokenType
		expectedLiteral string
	}{
		{LParen, "("},
		{Symbol, "defun"},
		{Symbol, "variable-with-dash"},
		{Int, "10"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)

	}
}

func TestFloatDeclaration(t *testing.T) {
	input := `(defun x 10.1)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LParen, "("},
		{Symbol, "defun"},
		{Symbol, "x"},
		{Float, "10.1"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)

	}
}

func TestPlusExprAreRecognized(t *testing.T) {
	input := `(+ 1 2)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LParen, "("},
		{Symbol, "+"},
		{Int, "1"},
		{Int, "2"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType)
		assert.Equal(t, tok.Literal, tt.expectedLiteral)
	}
}

func TestStringDeclaration(t *testing.T) {
	input := `("a" 'b' 'c" ab'cd x)`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LParen, "("},
		{String, `"a"`},
		{String, `'b'`},
		{IllegalToken, `'c"`},
		{IllegalToken, `ab'cd`},
		{Symbol, "x"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type, fmt.Sprintf("Token literal: %q", tok.Literal))
		assert.Equal(t, tt.expectedLiteral, tok.Literal)
	}
}

func TestValidateListOfStrings(t *testing.T) {
	input := `("a" "b")`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LParen, "("},
		{String, `"a"`},
		{String, `"b"`},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type, fmt.Sprintf("Token literal: %q", tok.Literal))
		assert.Equal(t, tt.expectedLiteral, tok.Literal)
	}

}

func TestValidateString(t *testing.T) {
	expr := `ab'cd`
	result := isValidSymbol(expr)
	assert.False(t, result)
}
