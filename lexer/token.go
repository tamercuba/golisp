package lexer

type TokenType uint8

const (
	LParen TokenType = iota
	RParen
	ReservedExpr
	Expr
	IllegalToken
	Int
	Float
)

type Token struct {
	Type    TokenType
	Literal string
	// TODO: Keep track of ch/col of the token, so we can have better error messages
}

func NewToken(ch byte, tokenType TokenType) Token {
	return Token{Literal: string(ch), Type: tokenType}
}

var reservedExpressions = map[string]TokenType{
	"let":    ReservedExpr,
	"lambda": ReservedExpr,
	"defun":  ReservedExpr,
	"+":      ReservedExpr,
	"-":      ReservedExpr,
	"*":      ReservedExpr,
	"/":      ReservedExpr,
}

func SetExprType(t *Token) {
	if _, ok := reservedExpressions[t.Literal]; ok {
		t.Type = ReservedExpr
	} else {
		t.Type = Expr
	}
}
