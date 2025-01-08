package lexer

import "github.com/tamercuba/golisp/token"

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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '(':
		tok = token.NewToken(l.ch, token.LParen)
	case ')':
		tok = token.NewToken(l.ch, token.RParen)
	case ' ':
		l.readChar()
		return l.NextToken()
	default:
		if isValidLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.Expr
		} else {
			tok = token.NewToken(l.ch, token.IllegalToken)
		}
		return tok
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isValidLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func isValidLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '-' || '0' <= ch && ch <= '9'
}
