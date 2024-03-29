package parser

import (
	"fmt"
	"strconv"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.INT_TYPE, token.FLOAT_TYPE, token.CHAR_TYPE, token.STRING_TYPE, token.DICT_TYPE:
		return p.parseDeclarationStatement()
	case token.IDENT:
		if p.nextTokenIs(token.ASSIGN) {
			return p.parseAssignmentStatement()
		} else {
			return p.parseExpressionStatement()
		}
	case token.RETURN:
		return p.parseReturnStatement()
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

	if p.nextTokenIs(token.LPAREN) {
		return p.parseFunctionDeclarationStatement()
	} else if p.nextTokenIs(token.LBRACKET) {
		return p.parseArrayDeclarationStatement()
	} else {
		return p.parseVariableDeclarationStatement()
	}
}

func (p *Parser) parseFunctionDeclarationStatement() ast.Statement {
	stmt := &ast.FunctionDeclarationStatement{}

	funcExp := &ast.FunctionExpression{
		Identifier: ast.Identifier{
			Name:        p.curToken.Literal,
			Type:        p.prevToken.Type,
			TypeLiteral: p.prevToken.Literal,
		},
	}

	p.advanceToken() // '('
	p.advanceToken() // params

	params := []*ast.Identifier{}
	for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) {
		if !token.IsDataType(p.curToken.Literal) {
			msg := fmt.Sprintf("expected data type, got= %s[%s]", p.curToken.Literal, p.curToken.Type)
			p.appendError(msg)
			return nil
		}
		param := &ast.Identifier{
			Type:        p.curToken.Type,
			TypeLiteral: p.curToken.Literal,
		}

		if !p.nextTokenIs(token.IDENT) {
			msg := fmt.Sprintf("expected IDENT after: %s, got= %s", p.curToken.Literal, p.nextToken.Type)
			p.appendError(msg)
			return nil
		}
		p.advanceToken() // IDENT
		param.Name = p.curToken.Literal

		params = append(params, param)

		if p.nextTokenIs(token.COMMA) {
			p.advanceToken() // ','
		}

		p.advanceToken()
	}
	funcExp.Parameters = params

	p.advanceToken() // '{'
	if !p.curTokenIs(token.LBRACE) {
		msg := fmt.Sprintf("expected '{' in function '%s' declaration, got=%s", funcExp.Identifier.Name, p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	funcExp.Body = p.parseBlockStatement()
	if funcExp.Body == nil {
		return nil
	}

	stmt.Function = funcExp

	return stmt
}

func (p *Parser) parseArrayDeclarationStatement() ast.Statement {
	stmt := &ast.ArrayDeclarationStatement{}
	stmt.Identifier = ast.Identifier{
		Name:        p.curToken.Literal,
		Type:        p.prevToken.Type,
		TypeLiteral: p.prevToken.Literal,
	}

	p.advanceToken() // '['

	if p.nextTokenIs(token.INT_VALUE) {
		p.advanceToken() // array size: INT_VALUE

		intVal, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
		if err != nil {
			msg := fmt.Sprintf("expected integer for array size, got %s", p.curToken.Literal)
			p.appendError(msg)
			return nil
		}

		stmt.Size = int(intVal)
	}

	if !p.nextTokenIs(token.RBRACKET) {
		msg := fmt.Sprintf("expected ']', got %s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}
	p.advanceToken() // ']'

	if p.nextTokenIs(token.SEMICOLON) { // did not initialize array
		p.advanceToken() // ';'
		return stmt
	}

	if !p.nextTokenIs(token.ASSIGN) || p.nextTokenIs(token.SEMICOLON) { // array initialization?
		msg := fmt.Sprintf("expected '=' or ';', got %s", p.nextToken.Literal)
		p.appendError(msg)
		return nil
	}

	p.advanceToken() // '='
	p.advanceToken() // expression

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

func (p *Parser) parseVariableDeclarationStatement() ast.Statement {
	stmt := &ast.VariableDeclarationStatement{}
	stmt.Identifier = ast.Identifier{
		Name:        p.curToken.Literal,
		Type:        p.prevToken.Type,
		TypeLiteral: p.prevToken.Literal,
	}

	if p.nextTokenIs(token.SEMICOLON) { // did not initialize variable
		p.advanceToken() // ';'
		return stmt
	}

	p.advanceToken() // '='
	p.advanceToken() // expression

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

func (p *Parser) parseAssignmentStatement() ast.Statement {
	stmt := &ast.AssignmentStatement{}
	stmt.Identifier = ast.Identifier{
		Name: p.curToken.Literal,
	}

	p.advanceToken() // '='
	p.advanceToken() // expression

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

	if !isIfExp && !isFuncExp {
		p.advanceToken() // ';'
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{}

	p.advanceToken() // expression

	stmt.ReturnValue = p.parseExpression(LOWEST)
	if stmt.ReturnValue == nil {
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
