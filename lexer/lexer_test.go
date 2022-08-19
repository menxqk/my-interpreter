package lexer

import (
	"testing"

	"github.com/menxqk/my-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `| abc int float char string dict
	10 35.50 'c' "A string"
	+ - / * ! = == != < <= > >= 
	, ; : ( ) [ ] { }
	`

	tests := []struct {
		Type    string
		Literal string
	}{
		{token.ILLEGAL, "|"},
		{token.IDENT, "abc"},
		{token.INT_TYPE, "int"},
		{token.FLOAT_TYPE, "float"},
		{token.CHAR_TYPE, "char"},
		{token.STRING_TYPE, "string"},
		{token.DICT_TYPE, "dict"},

		{token.INT_VALUE, "10"},
		{token.FLOAT_VALUE, "35.50"},
		{token.CHAR_VALUE, "c"},
		{token.STRING_VALUE, "A string"},

		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.ASSIGN, "="},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.LT, "<"},
		{token.LTE, "<="},
		{token.GT, ">"},
		{token.GTE, ">="},

		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		{token.EOF, "EOF"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("expected type %s, got=%s", tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("expected literal %q, got=%q", tt.Literal, tok.Literal)
		}
	}

}

func TestIllegalFloat(t *testing.T) {
	input := `
	10.0.00
	`
	tests := []struct {
		Type    string
		Literal string
	}{
		{token.ILLEGAL, "10.0.00"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("expected type %s, got=%s", tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("expected literal %q, got=%q", tt.Literal, tok.Literal)
		}
	}
}

func TestIllegalChar(t *testing.T) {
	input := `
	'cc' 'c
	`
	tests := []struct {
		Type    string
		Literal string
	}{
		{token.ILLEGAL, "'cc'"},
		{token.ILLEGAL, "'c"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("expected type %s, got=%s", tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("expected literal %q, got=%q", tt.Literal, tok.Literal)
		}
	}
}

func TestIllegalString(t *testing.T) {
	input := `"A str`
	tests := []struct {
		Type    string
		Literal string
	}{
		{token.ILLEGAL, "\"A str"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("expected type %s, got=%s", tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("expected literal %q, got=%q", tt.Literal, tok.Literal)
		}
	}
}
