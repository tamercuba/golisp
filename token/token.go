package token

type TokenType uint8

const (
	LParen TokenType = iota
	RParen
	Expr
	IllegalToken
)

type Token struct {
	Type    TokenType
	Literal string
	// TODO: Keep track of ch/col of the token, so we can have better error messages
}

func NewToken(ch byte, tokenType TokenType) Token {
	return Token{Literal: string(ch), Type: tokenType}
}
