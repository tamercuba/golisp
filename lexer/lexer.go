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
	default:
		tok = token.NewToken(l.ch, token.Atom)
	}

	l.readChar()
	return tok
}
