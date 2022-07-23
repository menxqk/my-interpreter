package parser

import (
	"fmt"

	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/lexer"
	"github.com/menxqk/my-interpreter/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[string]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.LTE:      LESSGREATER,
	token.GT:       LESSGREATER,
	token.GTE:      LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

type PrefixParseFn func() ast.Expression
type InfixParseFn func(ast.Expression) ast.Expression

type Parser struct {
	l *lexer.Lexer

	debug bool

	errors []string

	curToken  token.Token
	nextToken token.Token
	prevToken token.Token

	prefixParseFns map[string]PrefixParseFn
	infixParseFns  map[string]InfixParseFn
}

func New(l *lexer.Lexer, debug ...bool) *Parser {
	var d bool

	if len(debug) > 0 {
		d = debug[0]
	}

	p := &Parser{l: l, errors: make([]string, 0), debug: d}
	p.advanceToken()
	p.advanceToken()

	p.prefixParseFns = make(map[string]PrefixParseFn)
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixParseFn(token.INT_VALUE, p.parseIntegerLiteral)
	p.registerPrefixParseFn(token.FLOAT_VALUE, p.parseFloatLiteral)
	p.registerPrefixParseFn(token.CHAR_VALUE, p.parseCharLiteral)
	p.registerPrefixParseFn(token.STRING_VALUE, p.parseStringLiteral)
	p.registerPrefixParseFn(token.BANG, p.parsePrefixExpression)
	p.registerPrefixParseFn(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixParseFn(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefixParseFn(token.IF, p.parseIfExpression)

	p.infixParseFns = make(map[string]InfixParseFn)
	p.registerInfixParseFn(token.PLUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.MINUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParseFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixParseFn(token.EQ, p.parseInfixExpression)
	p.registerInfixParseFn(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfixParseFn(token.LT, p.parseInfixExpression)
	p.registerInfixParseFn(token.LTE, p.parseInfixExpression)
	p.registerInfixParseFn(token.GT, p.parseInfixExpression)
	p.registerInfixParseFn(token.GTE, p.parseInfixExpression)
	p.registerInfixParseFn(token.LPAREN, p.parseCallExpression)

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)

			// DEBUG INFO
			if p.debug {
				fmt.Printf("statement: %s\n", stmt.DebugString())
			}
		}

		p.advanceToken()
	}

	return program
}

func (p *Parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) appendError(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) advanceToken() {
	p.prevToken = p.curToken
	p.curToken = p.nextToken
	if !p.nextTokenIs(token.EOF) {
		p.nextToken = p.l.NextToken()
	}
}

func (p *Parser) curTokenIs(tokenType string) bool {
	return p.curToken.Type == tokenType
}

func (p *Parser) nextTokenIs(tokenType string) bool {
	return p.nextToken.Type == tokenType
}

func (p *Parser) prevTokenIs(tokenType string) bool {
	return p.prevToken.Type == tokenType
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextPrecedence() int {
	if p, ok := precedences[p.nextToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefixParseFn(tokenType string, fn PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfixParseFn(tokenType string, fn InfixParseFn) {
	p.infixParseFns[tokenType] = fn
}
