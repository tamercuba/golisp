package parser

import (
	"strconv"

	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

type Parser struct {
	l *lx.Lexer

	curToken  lx.Token
	peekToken lx.Token
}

func newParser(l *lx.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func ParseProgram(l *lx.Lexer) (*ast.Program, error) {
	p := newParser(l)
	program := &ast.Program{}
	program.ListStatements = []ast.ListExpression{}

	for p.curToken.Type != lx.EOF {
		if p.curToken.Type != lx.LParen {
			// Throw Error
		}

		list := p.parseList()
		if list != nil {
			program.ListStatements = append(program.ListStatements, *list)
		}
		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseList() *ast.ListExpression {
	list := ast.ListExpression{Token: p.curToken}
	for p.curToken.Type != lx.RParen {
		if p.curToken.Type == lx.EOF {
			// Throw Error
		}

		p.nextToken()

		switch p.curToken.Type {
		case lx.LParen:
			list.Elements = append(list.Elements, p.parseList())
		case lx.RParen:
			return &list
		case lx.ReservedExpr:
			list.Elements = append(list.Elements, *p.parseReservedExpr())
		case lx.Expr:
			list.Elements = append(list.Elements, *p.parseExpr())
		case lx.EOF:
			// Throw Error Not balanced parenthesis
		case lx.IllegalToken:
			// Throw Error Illegal character
		case lx.Int:
			value, err := strconv.Atoi(p.curToken.Literal)
			if err != nil {
				// Throw Error
			}
			list.Elements = append(list.Elements, &ast.IntLiteral{Token: p.curToken, Value: int32(value)})
		case lx.Float:
			value, err := strconv.ParseFloat(p.curToken.Literal, 64)
			if err != nil {
				// Throw Error
			}

			list.Elements = append(list.Elements, &ast.FloatLiteral{Token: p.curToken, Value: value})
		}
	}

	return nil
}

func (p *Parser) parseReservedExpr() *ast.Expression {
	return nil
}

func (p *Parser) parseExpr() *ast.Expression {
	return nil
}
