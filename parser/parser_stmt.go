package parser

import (
	"fmt"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {

	default:
		return p.parseExpressionStatement()
	}

}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{}

	stmt.Expression = p.parseExpression(LOWEST)
	if stmt.Expression == nil {
		return nil
	}

	if !p.nextTokenIs(token.SEMICOLON) {
		msg := fmt.Sprintf("missing ';' after %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ';'

	return stmt
}
