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
		if p.curToken.Type == lx.LParen {
			list := p.parseList()
			if list != nil {
				program.ListStatements = append(program.ListStatements, *list)
			}
		} else {
			p.nextToken()
		}
	}

	return program, nil
}

func (p *Parser) parseList() *ast.ListExpression {
	list := ast.ListExpression{Token: p.curToken}

	p.nextToken()

	if p.curToken.Type == lx.ReservedExpr {
		callExpr := p.parseReservedExpr()
		list.Elements = append(list.Elements, callExpr)
		return &list
	}

	for {
		switch p.curToken.Type {
		case lx.LParen:
			nestedList := p.parseList()
			list.Elements = append(list.Elements, nestedList)

		case lx.RParen:
			return &list

		case lx.Int:
			intLiteral := p.parseInt()
			list.Elements = append(list.Elements, intLiteral)
			p.nextToken()
		case lx.Float:
			floatLiteral := p.parseFloat()
			list.Elements = append(list.Elements, floatLiteral)
			p.nextToken()
		case lx.ReservedExpr:
			reservedExpr := p.parseReservedExpr()
			list.Elements = append(list.Elements, reservedExpr)
		case lx.EOF:
			panic("Unbalanced parentheses: EOF reached while parsing a list")
		default:
			panic(fmt.Sprintf("Unexpected token in list: %v", p.curToken))
		}
	}
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
	callExpr := &ast.CallExpression{Token: p.curToken, Arguments: []ast.Node{}}
	p.nextToken()

	for {
		switch p.curToken.Type {
		case lx.LParen:
			list := p.parseList()
			callExpr.Arguments = append(callExpr.Arguments, list)
		case lx.RParen:
			return callExpr
		case lx.Int:
			intLiteral := p.parseInt()
			callExpr.Arguments = append(callExpr.Arguments, intLiteral)
			p.nextToken()
		case lx.Float:
			floatLiteral := p.parseFloat()
			callExpr.Arguments = append(callExpr.Arguments, floatLiteral)
			p.nextToken()
		case lx.ReservedExpr:
			nestedCall := p.parseReservedExpr()
			callExpr.Arguments = append(callExpr.Arguments, nestedCall)
		case lx.EOF:
			panic("Unbalanced parentheses: EOF reached while parsing a reserved expression")
		default:
			panic(fmt.Sprintf("Unexpected token in reserved expression: %v", p.curToken))
		}
	}
}
func (p *Parser) parseExpr() *ast.CallExpression {
	fmt.Print("ERROR PARSE EXPR\n")
	return nil
}
