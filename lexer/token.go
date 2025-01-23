package lexer

import "fmt"

type TokenType uint8

const (
	LParen TokenType = iota
	RParen
	Symbol
	IllegalToken
	String
	Int
	Float
	Bool
	Void
	EOF
)

type Position struct {
	Ch  int
	Col int
}

func (p Position) String() string {
	return fmt.Sprintf("%d:%d ", p.Col, p.Ch)
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}

func NewPos(ch int, col int) *Position {
	return &Position{Ch: ch, Col: col}
}

func NewToken(char byte, tokenType TokenType, ch int, col int) Token {
	return Token{Literal: string(char), Type: tokenType, Pos: *NewPos(ch, col)}
}

func (t *Token) SetPos(ch int, col int) {
	t.Pos.Col = col
	t.Pos.Ch = ch

}
