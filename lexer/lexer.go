package lexer

import "regexp"

type Lexer struct {
	input   string
	pos     int // points to current char
	readPos int // points to after current char
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Only works with ASCII chars
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '(':
		tok = NewToken(l.ch, LParen)
	case ')':
		tok = NewToken(l.ch, RParen)
	case ' ':
		l.readChar()
		return l.NextToken()
	default:
		if !isValidChar(l.ch) {
			tok = NewToken(l.ch, IllegalToken)
			l.readChar()
			return tok
		}

		tok.Literal = l.readExpr()
		if isValidNumber(tok.Literal) {
			if isInteger(tok.Literal) {
				tok.Type = Int
			} else {
				tok.Type = Float
			}
			return tok
		} else if isValidExpr(tok.Literal) {
			SetExprType(&tok)
			return tok
		} else {
			tok.Type = IllegalToken
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readExpr() string {
	pos := l.pos
	for isValidChar(l.ch) || l.ch == '.' || l.ch == '/' || l.ch == '-' {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func isValidNumber(expr string) bool {
	re := regexp.MustCompile(`^\d+(\.\d+)?$`)
	return re.MatchString(expr)
}

func isInteger(expr string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(expr)
}

func isValidExpr(expr string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9+\-*/]+$`)
	return re.MatchString(expr)
}

func isValidChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch == '+' || ch == '-' || ch == '*' || ch == '/'
}
