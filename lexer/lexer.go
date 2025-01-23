package lexer

import (
	"regexp"
)

type Lexer struct {
	input   string
	pos     int // points to current char
	readPos int // points to after current char
	ch      byte
	posCol  int
	posCh   int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, posCol: 0, posCh: -1}
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

	if l.ch == '\n' {
		l.posCol++
		l.posCh = -1
	} else {
		l.posCh++
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '(':
		tok = NewToken(l.ch, LParen, l.posCh, l.posCol)
	case ')':
		tok = NewToken(l.ch, RParen, l.posCh, l.posCol)
	case '\n', ' ':
		l.readChar()
		return l.NextToken()
	case 0:
		tok = NewToken(0, EOF, l.posCh, l.posCol)
		return tok
	default:
		if !isValidChar(l.ch) {
			tok = NewToken(l.ch, IllegalToken, l.posCh, l.posCol)
			l.readChar()
			return tok
		}
		tok.SetPos(l.posCh, l.posCol)
		tok.Literal = l.readExpr()
		if isValidNumber(tok.Literal) {
			if isInteger(tok.Literal) {
				tok.Type = Int
			} else {
				tok.Type = Float
			}
			return tok
		} else if isValidString(tok.Literal) {
			tok.Type = String
			return tok
		} else if isValidBool(tok.Literal) {
			tok.Type = Bool
			return tok
		} else if isValidNil(tok.Literal) {
			tok.Type = Void
			return tok
		} else if isValidSymbol(tok.Literal) {
			tok.Type = Symbol
			return tok
		} else {
			tok.Type = IllegalToken
			return tok
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

func validateRegex(expr string, regex string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(expr)
}

func isValidNil(expr string) bool {
	return validateRegex(expr, `^(nil)$`)
}

func isValidString(expr string) bool {
	return validateRegex(expr, `^'[^']*'$`) || validateRegex(expr, `^"[^"]*"$`)
}

func isValidNumber(expr string) bool {
	return validateRegex(expr, `^\d+(\.\d+)?$`)
}

func isInteger(expr string) bool {
	return validateRegex(expr, `^\d+$`)
}

func isValidSymbol(expr string) bool {
	return validateRegex(expr, `^[a-zA-Z0-9+\-*/^]+$`)
}

func isValidBool(expr string) bool {
	return validateRegex(expr, `^(true|false)$`)
}

func isValidChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '"' || ch == '\''
}
