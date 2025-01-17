package lexer

import "testing"

func TestDefunBasicDeclaration(t *testing.T) {
	input := `(defun x abc)`

	tests := []struct {
		expectedType    TokenType
		exceptedLiteral string
		posCh           int
		posCol          int
	}{
		{LParen, "(", 0, 0},
		{ReservedExpr, "defun", 1, 0},
		{Expr, "x", 7, 0},
		{Expr, "abc", 9, 0},
		{RParen, ")", 12, 0},
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
		if tok.Pos.Ch != tt.posCh {
			t.Fatalf("TestDefunBasicDeclaration[%d] - char [%q] position is wrong, expected=%d, got%d", i, tok.Literal, tt.posCh, tok.Pos.Ch)
		}
		if tok.Pos.Col != tt.posCol {
			t.Fatalf("TestDefunBasicDeclaration[%d] - char line is wrong, expected=%d, got%d", i, tt.posCol, tok.Pos.Col)
		}
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
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Literal != tt.literal {
			t.Fatalf("TestPositionTracker([%d] - Token literal is wrong, expected=%q, got=%q", i, tt.literal, tok.Literal)
		}
		if tok.Pos.Ch != tt.ch {
			t.Fatalf("TestPositionTracker([%d] - Token [%q] position is wrong, expected=%d, got=%d", i, tok.Literal, tt.ch, tok.Pos.Ch)
		}
		if tok.Pos.Col != tt.col {
			t.Fatalf("TestPositionTracker([%d] - Token line [%q] position is wrong, expected=%d, got=%d", i, tok.Literal, tt.col, tok.Pos.Col)
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
