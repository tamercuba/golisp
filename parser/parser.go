package parser

import (
	"fmt"
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
			fmt.Printf("ERROR PARSE LIST ) MISSING 1")
			// Throw Error
		}

		p.nextToken()

		switch p.curToken.Type {
		case lx.LParen:
			list.Elements = append(list.Elements, p.parseList())
		case lx.RParen:
			return &list
		case lx.ReservedExpr:
			list.Elements = append(list.Elements, p.parseReservedExpr())
			return &list
		case lx.Expr:
			list.Elements = append(list.Elements, p.parseExpr())
		case lx.EOF:
			// Throw Error Not balanced parenthesis
			panic(p.peekToken)
		case lx.IllegalToken:
			// Throw Error Illegal character
			panic(lx.IllegalToken)
		case lx.Int:
			list.Elements = append(list.Elements, p.parseInt())
		case lx.Float:
			list.Elements = append(list.Elements, p.parseFloat())
		}
	}

	return nil
}

func (p *Parser) parseInt() *ast.IntLiteral {
	value, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		fmt.Printf("ERROR PARSE LIST INT VALUE ATOI")
		// Throw Error
	}

	return &ast.IntLiteral{Token: p.curToken, Value: int32(value)}
}

func (p *Parser) parseFloat() *ast.FloatLiteral {
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		fmt.Printf("ERROR PARSE LIST FLOAT VALUE ATOI")
		// Throw Error
	}
	return &ast.FloatLiteral{Token: p.curToken, Value: value}
}

func (p *Parser) parseReservedExpr() *ast.CallExpression {
	callExpr := ast.CallExpression{Token: p.curToken, Arguments: []ast.Expression{}}

	for p.curToken.Type != lx.RParen {
		p.nextToken()
		if p.curToken.Type == lx.EOF {
			fmt.Print("ERROR EOF\n")
			// Throw Error
		}

		switch p.curToken.Type {
		case lx.LParen:
			list := p.parseList()
			if list == nil {
				// Throw Error
				fmt.Print("ERROR LPAREN LIST == NIL\n")
			}
			callExpr.Arguments = append(callExpr.Arguments, list)
		case lx.RParen:
			return &callExpr
		case lx.IllegalToken:
			// Thor Error
		case lx.Int:
			callExpr.Arguments = append(callExpr.Arguments, p.parseInt())
		case lx.Float:
			callExpr.Arguments = append(callExpr.Arguments, p.parseFloat())
		}

	}
	// Maybe throw an error?
	fmt.Print("ERROR OUT OF SWITCH\n")
	return nil
}

func (p *Parser) parseExpr() *ast.CallExpression {
	fmt.Print("ERROR PARSE EXPR\n")
	return nil
}
