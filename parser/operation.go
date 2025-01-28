package parser

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseOperation() *ast.OperationNode {
	//  c p
	// (+ 1 2)

	if p.curToken.Type != lx.Symbol {
		panic(fmt.Sprintf("%v Type Error. %v isnt a valid function name", p.peekToken.Pos, p.peekToken))

	}
	token := p.curToken
	name := p.parseSymbol()
	params := []ast.Node{}

	for p.curToken.Type != lx.RParen {
		switch p.curToken.Type {
		case lx.LParen:
			l := p.parseList()
			if l != nil {
				params = append(params, l)
			}
		case lx.Float:
			params = append(params, p.parseFloat())
		case lx.Int:
			params = append(params, p.parseInt())
		case lx.Symbol:
			params = append(params, p.parseSymbol())
		case lx.String:
			params = append(params, p.parseString())
		case lx.Bool:
			params = append(params, p.parseBoolean())
		case lx.Void:
			params = append(params, p.parseVoid())
		case lx.EOF:
			panic(fmt.Sprintf("%v( not closed, expect ).", token.Pos))
		default:
			panic(fmt.Sprintf("Unexpected token in list: %+v", p.curToken))
		}
	}

	return ast.NewOperationNode(token, name, params)
}
