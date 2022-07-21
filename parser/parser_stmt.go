package parser

import (
	"fmt"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.INT_TYPE, token.FLOAT_TYPE, token.CHAR_TYPE, token.STRING_TYPE:
		// declaration statement
		return p.parseDeclarationStatement()
	case token.IDENT:
		if p.nextTokenIs(token.ASSIGN) {
			return p.parseAssignmentStatement()
		} else {
			return p.parseExpressionStatement()
		}
	default:
		return p.parseExpressionStatement()
	}

}

func (p *Parser) parseDeclarationStatement() ast.Statement {
	p.advanceToken() // ident
	if !p.curTokenIs(token.IDENT) {
		msg := fmt.Sprintf("expected identifier after: %s; got=%s", p.prevToken.Literal, p.curToken.Type)
		p.appendError(msg)
		return nil
	}

	if p.nextTokenIs(token.ASSIGN) {
		return p.parseVariableDeclarationStatement()
	} else if p.nextTokenIs(token.LPAREN) {
		return p.parseFunctionDeclarationStatement()
	} else {
		msg := fmt.Sprintf("expected '(' or '=' after: %s %s", p.prevToken.Literal, p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
}

func (p *Parser) parseVariableDeclarationStatement() ast.Statement {
	// TODO
	return nil
}

func (p *Parser) parseFunctionDeclarationStatement() ast.Statement {
	// TODO
	return nil
}

func (p *Parser) parseAssignmentStatement() ast.Statement {
	// TODO
	return nil
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{}

	stmt.Expression = p.parseExpression(LOWEST)
	if stmt.Expression == nil {
		return nil
	}

	_, isIfExp := stmt.Expression.(*ast.IfExpression)
	_, isFuncExp := stmt.Expression.(*ast.FunctionExpression)
	if !p.nextTokenIs(token.SEMICOLON) && !isIfExp && !isFuncExp {
		msg := fmt.Sprintf("missing ';' after %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ';'

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{}
	block.Statements = []ast.Statement{}

	p.advanceToken() // after '{'

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		p.advanceToken()
		if stmt == nil {
			return nil
		}
		block.Statements = append(block.Statements, stmt)
	}

	return block
}
