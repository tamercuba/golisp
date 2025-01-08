package lexer

import (
	"github.com/tamercuba/golisp/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(defun x 10)`

	tests := []struct {
		expectedType    token.TokenType
		exceptedLiteral string
	}{
		{token.LParen, "("},
		{token.Atom, "d"},
		{token.Atom, "e"},
		{token.Atom, "f"},
		{token.Atom, "u"},
		{token.Atom, "n"},
		{token.Atom, " "},
		{token.Atom, "x"},
		{token.Atom, " "},
		{token.Atom, "1"},
		{token.Atom, "0"},
		{token.RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType is wrong, expected=%q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("tests[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}
