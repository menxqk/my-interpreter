package parser

import (
	"github.com/menxqk/my-interpreter/ast"
	"github.com/menxqk/my-interpreter/lexer"
	"github.com/menxqk/my-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	nextToken token.Token
	prevToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: make([]string, 0)}
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for tok := p.l.NextToken(); tok.Type != token.EOF; tok = p.l.NextToken() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advanceToken()
	}

	return program
}

func (p *Parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *Parser) Errors() []string {
	return p.Errors()
}

func (p *Parser) advanceToken() {
	p.prevToken = p.curToken
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
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
