package lexer

import "testing"

func TestDefunBasicDeclaration(t *testing.T) {
	input := `(defun x abc)`

	tests := []struct {
		expectedType    TokenType
		exceptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "defun"},
		{Expr, "x"},
		{Expr, "abc"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestDefunBasicDeclaration[%d] - ype is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestDefunBasicDeclaration[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}

func TestDefunDeclarationWithDashChar(t *testing.T) {
	input := `(defun variable-with-dash 10)`

	tests := []struct {
		expectedType    TokenType
		exceptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "defun"},
		{Expr, "variable-with-dash"},
		{Int, "10"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestDefunDeclarationWithDashChar[%d] - ype is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestDefunDeclarationWithDashChar[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}

func TestFloatDeclaration(t *testing.T) {
	input := `(defun x 10.1)`

	tests := []struct {
		expectedType    TokenType
		exceptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "defun"},
		{Expr, "x"},
		{Float, "10.1"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestFloatDeclaration[%d] - ype is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestFloatDeclaration[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}

func TestPlusExprAreRecognized(t *testing.T) {
	input := `(+ 1 2)`

	tests := []struct {
		expectedType    TokenType
		exceptedLiteral string
	}{
		{LParen, "("},
		{ReservedExpr, "+"},
		{Int, "1"},
		{Int, "2"},
		{RParen, ")"},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestFloatDeclaration[%d] - ype is wrong, expected=%d, got=%d. Token: %q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("TestFloatDeclaration[%d] - literal is wrong, expected=%q, got%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}
