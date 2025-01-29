package parser

import (
	lx "github.com/tamercuba/golisp/lexer"
	"github.com/tamercuba/golisp/parser/ast"
)

func (p *Parser) parseOperation() (*ast.OperationNode, error) {
	//  c p
	// (+ 1 2)
	if p.curToken.Type != lx.Symbol {
		return nil, NewParseError("isn't a valid function name", p.peekToken)
	}

	token := p.curToken
	name := p.parseSymbol()
	params := []ast.Node{}

	for p.curToken.Type != lx.RParen {
		switch p.curToken.Type {
		case lx.LParen:
			l, err := p.parseList()
			if err != nil {
				return nil, err
			}
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
			return nil, NewParseError("not closed, expect )", token)
		default:
			return nil, NewParseError("isn't a valid list element", p.curToken)
		}
	}

	return ast.NewOperationNode(token, name, params), nil
}
