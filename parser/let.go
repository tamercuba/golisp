package parser

import (
	"fmt"

	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseLet() *ast.LetDeclaration {
	firstToken := p.curToken
	//   c  p
	// (let x 10)

	if p.peekToken.Type != lx.Symbol {
		panic(fmt.Sprintf("%q Type Error. %q isnt a valid binding name", p.peekToken.Pos, p.peekToken))
	}
	p.nextToken()
	bindingName := p.parseSymbol()

	//        c p
	// (let x 10)
	var bindingValue ast.Node
	switch p.curToken.Type {
	case lx.LParen:
		bindingValue = p.parseList()
	case lx.Symbol:
		bindingValue = p.parseSymbol()
	case lx.Int:
		bindingValue = p.parseInt()
	case lx.Float:
		bindingValue = p.parseFloat()
	case lx.String:
		bindingValue = p.parseString()
	default:
		panic(fmt.Sprintf("%q Type Error. %q isnt a valid binding value", p.peekToken.Pos, p.peekToken))
	}

	p.nextToken()

	return ast.NewLetDeclaration(firstToken, bindingName, bindingValue)
}
