package parser

import (
	"fmt"

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
	program.ListStatements = []ast.Node{}

	for p.curToken.Type != lx.EOF {
		switch p.curToken.Type {
		case lx.LParen:
			list := p.parseList()
			if list != nil {
				program.ListStatements = append(program.ListStatements, list)
			}
		case lx.Int:
			n := p.parseInt()
			program.ListStatements = append(program.ListStatements, n)
		case lx.Float:
			n := p.parseFloat()
			program.ListStatements = append(program.ListStatements, n)
		case lx.String:
			n := p.parseString()
			program.ListStatements = append(program.ListStatements, n)
		case lx.Bool:
			n := p.parseBoolean()
			program.ListStatements = append(program.ListStatements, n)
		case lx.Void:
			n := p.parseVoid()
			program.ListStatements = append(program.ListStatements, n)
		default:
			panic("INVALID SYNTAX")
		}
	}

	return program, nil
}

func (p *Parser) parseList() ast.Node {
	// Expect curToken.Type == LParen
	if p.peekToken.Type == lx.Symbol {
		result := p.parseNextSymbol()
		if result != nil {
			p.nextToken()
			return result
		}
	}

	list := ast.NewListExpression(p.curToken)
	p.nextToken()

	for {
		switch p.curToken.Type {
		case lx.LParen:
			nestedList := p.parseList()
			if nestedList != nil {
				list.Append(nestedList)
			}
		case lx.RParen:
			p.nextToken()
			return list
		case lx.Int:
			list.Append(p.parseInt())
		case lx.Float:
			list.Append(p.parseFloat())
		case lx.String:
			list.Append(p.parseString())
		case lx.Symbol:
			list.Append(p.parseSymbol())
		case lx.Bool:
			list.Append(p.parseBoolean())
		case lx.Void:
			list.Append(p.parseVoid())
		case lx.EOF:
			panic(fmt.Sprintf("%v( not closed, expect ).", list.GetToken().Pos))
		default:
			panic(fmt.Sprintf("Unexpected token in list: %+v", p.curToken))
		}
	}
}

func (p *Parser) parseBoolean() *ast.Boolean {
	result := ast.NewBoolean(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseInt() *ast.IntLiteral {
	result := ast.NewIntLiteral(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseFloat() *ast.FloatLiteral {
	result := ast.NewFloatLiteral(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseString() *ast.StringLiteral {
	result := ast.NewStringLiteral(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseVoid() *ast.VoidNode {
	result := ast.NewVoidNode(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseSymbol() *ast.Symbol {
	result := ast.NewSymbol(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseNextSymbol() ast.Node {
	for {
		switch p.peekToken.Literal {
		case "let", "define":
			p.nextToken()
			return p.parseVar()
		case "lambda":
			p.nextToken()
			return p.parseLambda()
		default:
			// Nothing special
			return nil
		}
	}
}
