package parser

import (
	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseVar() (*ast.VarDifinitionNode, error) {
	firstToken := p.curToken
	//   c  p
	// (let x 10)

	if p.peekToken.Type != lx.Symbol {
		return nil, NewParseError("isn't a valid binding name", p.peekToken)
	}
	p.nextToken()
	bindingName := p.parseSymbol()

	//        c p
	// (let x 10)
	var bindingValue ast.Node
	var err error = nil

	switch p.curToken.Type {
	case lx.LParen:
		bindingValue, err = p.parseList()
		if err != nil {
			return nil, err
		}
	case lx.Symbol:
		bindingValue = p.parseSymbol()
	case lx.Int:
		bindingValue = p.parseInt()
	case lx.Float:
		bindingValue = p.parseFloat()
	case lx.String:
		bindingValue = p.parseString()
	case lx.Bool:
		bindingValue = p.parseBoolean()
	case lx.Void:
		bindingValue = p.parseVoid()
	default:
		return nil, NewParseError("isn't a valid binding value", p.peekToken)
	}

	return ast.NewVarDifinitionNode(firstToken, bindingName, bindingValue), nil
}
