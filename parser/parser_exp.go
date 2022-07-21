package parser

import (
	"fmt"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	if p.curTokenIs(token.ILLEGAL) {
		msg := fmt.Sprintf("illegal token: %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	prefixFn := p.prefixParseFns[p.curToken.Type]
	if prefixFn == nil {
		msg := fmt.Sprintf("no prefix parse function for: %s", p.curToken.Type)
		p.appendError(msg)
		return nil
	}

	leftExp := prefixFn()

	for !p.nextTokenIs(token.SEMICOLON) && precedence < p.nextPrecedence() {
		infix := p.infixParseFns[p.nextToken.Type]
		if infix == nil {
			return leftExp
		}
		p.advanceToken() // next operator

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	exp := &ast.Identifier{}
	exp.Name = p.curToken.Literal

	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{}
	exp.Operator = p.curToken.Literal

	p.advanceToken() // expression

	exp.Expression = p.parseExpression(PREFIX)
	if exp.Expression == nil {
		return nil
	}

	return exp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	exp := &ast.GroupedExpression{}

	p.advanceToken() // expression

	exp.Expression = p.parseExpression(LOWEST)
	if exp.Expression == nil {
		return nil
	}

	if !p.nextTokenIs(token.RPAREN) {
		msg := fmt.Sprintf("missing ')' after: %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ')'

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{}

	exp.Left = left
	exp.Operator = p.curToken.Literal

	precedence := p.curPrecedence()
	p.advanceToken() // right expression
	exp.Right = p.parseExpression(precedence)
	if exp.Right == nil {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{}

	if !p.nextTokenIs(token.LPAREN) {
		msg := fmt.Sprintf("expected '(' got=%s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // '('
	p.advanceToken() // expression

	exp.Condition = p.parseExpression(LOWEST)
	if exp.Condition == nil {
		return nil
	}

	if !p.nextTokenIs(token.RPAREN) {
		msg := fmt.Sprintf("expected ')' got=%s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ')'

	if !p.nextTokenIs(token.LBRACE) {
		msg := fmt.Sprintf("expected '{' got=%s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // '{'

	exp.Consequence = p.parseBlockStatement()
	if exp.Consequence == nil {
		return nil
	}

	if p.nextTokenIs(token.ELSE) {
		p.advanceToken() // else
		if !p.nextTokenIs(token.LBRACE) {
			msg := fmt.Sprintf("expected '{' got=%s", p.nextToken.Literal)
			p.appendError(msg)
			return nil
		}
		p.advanceToken() // '{'

		exp.Alternative = p.parseBlockStatement()
		if exp.Alternative == nil {
			return nil
		}
	}

	return exp
}

func (p *Parser) parseCallExpression(left ast.Expression) ast.Expression {
	// TODO
	return nil
}
