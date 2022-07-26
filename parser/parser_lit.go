package parser

import (
	"fmt"
	"strconv"

	"github.com/menxqk/my-interpreter/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not convert %s to integer", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not convert %s to float", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseCharLiteral() ast.Expression {
	lit := &ast.CharLiteral{}

	if len(p.curToken.Literal) > 0 {
		lit.Value = rune(p.curToken.Literal[0])
	} else {
		msg := fmt.Sprintf("could not convert %s to char", p.curToken.Literal)
		p.appendError(msg)
		return nil
	}

	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	lit := &ast.StringLiteral{}

	lit.Value = p.curToken.Literal

	return lit
}

func (p *Parser) parseBoolean() ast.Expression {
	switch p.curToken.Literal {
	case "true":
		return &ast.BooleanLiteral{Value: true}
	case "false":
		return &ast.BooleanLiteral{Value: false}
	default:
		return nil
	}
}

func (p *Parser) parseNull() ast.Expression {
	return &ast.NullLiteral{}
}
