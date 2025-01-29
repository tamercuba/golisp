package parser

import (
	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseLambda() (*ast.LambdaNode, error) {
	//  c      p
	// (lambda (x y) (+ x y))
	firstToken := p.curToken
	p.nextToken()

	funcArgs, err := p.getFunctionArgs()
	if err != nil {
		return nil, err
	}

	body, err := p.getFunctionBody()
	if err != nil {
		return nil, err
	}

	return ast.NewLambdaNode(firstToken, funcArgs, body), nil
}

func (p *Parser) getFunctionBody() (ast.Node, error) {
	//               cp
	// (lambda (x y) (+ x y))
	if p.peekToken.Type != lx.LParen {
		return nil, NewParseError("isn't a valid function body, should be a list.", p.peekToken)
	}

	p.nextToken()
	body, err := p.parseList()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (p *Parser) getFunctionArgs() ([]ast.Symbol, error) {
	//          cp
	// (defun x (y) (+ 1 y))
	if p.curToken.Type != lx.LParen {
		return nil, NewParseError("isn't a valid function argument, should be a list of symbols", p.peekToken)
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
			return args, nil
		default:
			return nil, NewParseError("should be a valid function argument or )", p.curToken)
		}
	}
}
