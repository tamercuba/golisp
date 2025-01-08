package lexer

import (
	"github.com/tamercuba/golisp/token"
	"testing"
)

func TestDefunBasicDeclaration(t *testing.T) {
	input := `(defun x 10)`

	tests := []struct {
		expectedType    token.TokenType
		exceptedLiteral string
	}{
		{token.LParen, "("},
		{token.Expr, "defun"},
		{token.Expr, "x"},
		{token.Expr, "10"},
		{token.RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestDefunBasicDeclaration[%d] - tokenType is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestDefunBasicDeclaration[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}

func TestDefunDeclarationWithDashChar(t *testing.T) {
	input := `(defun variable-with-dash 10)`

	tests := []struct {
		expectedType    token.TokenType
		exceptedLiteral string
	}{
		{token.LParen, "("},
		{token.Expr, "defun"},
		{token.Expr, "variable-with-dash"},
		{token.Expr, "10"},
		{token.RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestDefunDeclarationWithDashChar[%d] - tokenType is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestDefunDeclarationWithDashChar[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}
