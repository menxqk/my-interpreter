package parser

import (
	"fmt"
	"strconv"

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
		msg := fmt.Sprintf("expected '(' got= %s", p.nextToken.Literal)
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
		msg := fmt.Sprintf("expected ')' got= %s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ')'

	if !p.nextTokenIs(token.LBRACE) {
		msg := fmt.Sprintf("expected '{' got= %s", p.nextToken.Literal)
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
			msg := fmt.Sprintf("expected '{' got= %s", p.nextToken.Literal)
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
	exp := &ast.CallExpression{}
	exp.Identifier = ast.Identifier{
		Name: left.Literal(),
	}

	p.advanceToken() // call arguments
	args := []ast.Expression{}
	for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) {
		exp := p.parseExpression(LOWEST)
		if exp == nil {
			return nil
		}
		args = append(args, exp)

		if p.nextTokenIs(token.COMMA) {
			p.advanceToken() // ','
		}

		p.advanceToken()
	}
	exp.Arguments = args

	return exp
}

func (p *Parser) parseCollectionElementExpression(left ast.Expression) ast.Expression {
	switch p.nextToken.Type {
	case token.INT_VALUE:
		return p.parseArrayElementExpression(left)
	case token.STRING_VALUE:
		return p.parseDictElementExpression(left)
	default:
		return nil
	}
}

func (p *Parser) parseArrayElementExpression(left ast.Expression) ast.Expression {
	exp := &ast.ArrayElementExpression{}
	exp.Identifier = ast.Identifier{
		Name: left.Literal(),
	}

	p.advanceToken() // array element index
	if !p.curTokenIs(token.INT_VALUE) {
		msg := fmt.Sprintf("expected integer value for array element index, got %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	val, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse integer for array element index: %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	exp.Index = int(val)

	if !p.nextTokenIs(token.RBRACKET) {
		msg := fmt.Sprintf("expected ']' after %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ']'

	if p.nextTokenIs(token.ASSIGN) {
		p.advanceToken() // '='
		p.advanceToken() // expression

		elemExp := p.parseExpression(LOWEST)
		if exp == nil {
			return nil
		}
		exp.Expression = elemExp
	}

	return exp
}

func (p *Parser) parseDictElementExpression(left ast.Expression) ast.Expression {
	exp := &ast.DictElementExpression{}
	exp.Identifier = ast.Identifier{
		Name: left.Literal(),
	}

	p.advanceToken() // dict key
	if !p.curTokenIs(token.STRING_VALUE) {
		msg := fmt.Sprintf("expected string for dict key, got %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	exp.Key = p.curToken.Literal

	if !p.nextTokenIs(token.RBRACKET) {
		msg := fmt.Sprintf("expected ']' after %s", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ']'

	if p.nextTokenIs(token.ASSIGN) {
		p.advanceToken() // '='
		p.advanceToken() // expression

		elemExp := p.parseExpression(LOWEST)
		if exp == nil {
			return nil
		}
		exp.Expression = elemExp
	}

	return exp
}
