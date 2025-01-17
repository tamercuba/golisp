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

type Position struct {
	Ch  int
	Col int
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
	// TODO: Keep track of ch/col of the token, so we can have better error messages
}

func newPos(ch int, col int) *Position {
	return &Position{Ch: ch, Col: col}
}

func NewToken(char byte, tokenType TokenType, ch int, col int) Token {
	return Token{Literal: string(char), Type: tokenType, Pos: *newPos(ch, col)}
}

func (t *Token) SetPos(ch int, col int) {
	t.Pos.Col = col
	t.Pos.Ch = ch

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
