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
		if p.curToken.Type == lx.LParen {
			list := p.parseList()
			if list != nil {
				program.ListStatements = append(program.ListStatements, list)
			}
		} else {
			// TODO: Maybe throw an error? Idk yet
			p.nextToken()
		}
	}

	return program, nil
}

func (p *Parser) parseList() ast.Node {
	// Expect curToken.Type == LParen
	if p.curToken.Type != lx.LParen {
		panic(fmt.Sprintf("[%q] Invalid syntax, %q given, '(' expected", p.curToken.Pos, p.curToken))
	}

	if p.peekToken.Type == lx.Symbol {
		result := p.parseNextSymbol()
		if result != nil {
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
			return list
		case lx.Int:
			list.Append(p.parseInt())
		case lx.Float:
			list.Append(p.parseFloat())
		case lx.String:
			list.Append(p.parseString())
		case lx.Symbol:
			list.Append(p.parseSymbol())
		case lx.EOF:
			panic("Unbalanced parentheses: EOF reached while parsing a list")
		default:
			panic(fmt.Sprintf("Unexpected token in list: %v", p.curToken))
		}
	}
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

func (p *Parser) parseSymbol() *ast.Symbol {
	result := ast.NewSymbol(p.curToken)
	p.nextToken()
	return result
}

func (p *Parser) parseNextSymbol() ast.Node {
	for {
		switch p.peekToken.Literal {
		case "defun":
			p.nextToken()
			return p.parseDefun()
		case "let":
			p.nextToken()
			return p.parseLet()
		default:
			// Nothing special
			return nil
		}
	}
}
