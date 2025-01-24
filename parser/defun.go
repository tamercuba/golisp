package parser

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseDefun() *ast.FunctionDeclaration {
	firstToken := p.curToken
	//    c   p
	// (defun x (y) (+ 1 y))
	if p.peekToken.Type != lx.Symbol {
		panic(fmt.Sprintf("%v Type Error. %v isnt a valid function name", p.peekToken.Pos, p.peekToken))
	}
	p.nextToken()
	funcName := ast.NewSymbol(p.curToken)
	p.nextToken()

	funcArgs := p.getFunctionArgs()
	body := p.getFunctionBody()
	return ast.NewFunctionDeclaration(firstToken, funcName, funcArgs, body)
}

func (p *Parser) parseLambda() *ast.LambdaNode {
	//  c      p
	// (lambda (x y) (+ x y))
	firstToken := p.curToken
	p.nextToken()
	funcArgs := p.getFunctionArgs()
	body := p.getFunctionBody()

	return ast.NewLambdaNode(firstToken, funcArgs, body)
}

func (p *Parser) getFunctionBody() ast.Node {
	//               cp
	// (lambda (x y) (+ x y))
	if p.peekToken.Type != lx.LParen {
		panic(fmt.Sprintf("%v Type Error. Function body should be a list, not %v", p.peekToken.Pos, p.peekToken))
	}
	p.nextToken()
	body := p.parseList()
	return body
}

func (p *Parser) getFunctionArgs() []ast.Symbol {
	//          cp
	// (defun x (y) (+ 1 y))
	if p.curToken.Type != lx.LParen {
		panic(fmt.Sprintf("%v Type Error. Function args should be a List, not %v", p.peekToken.Pos, p.peekToken))
	}

	p.nextToken()

	//           c p
	// (defun x (y z) (+ z y))
	args := []ast.Symbol{}
	for {
		switch p.curToken.Type {
		case lx.Symbol:
			newParam := ast.NewSymbol(p.curToken)
			args = append(args, *newParam)
			p.nextToken()
		case lx.RParen:
			return args
		default:
			panic(fmt.Sprintf("%+v Invalid Syntax. %+v Should be a valid function argument or )", p.curToken.Pos, p.curToken))
		}
	}
}
