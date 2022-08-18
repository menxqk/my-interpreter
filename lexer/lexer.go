package lexer

import (
	"fmt"
	"strings"

	"github.com/menxqk/my-interpreter/token"
)

type Lexer struct {
	input []rune

	curPos  int
	nextPos int

	char rune

	debug bool
}

func New(input string, debug ...bool) *Lexer {
	var d bool

	if len(debug) > 0 {
		d = debug[0]
	}

	l := &Lexer{input: []rune(input), debug: d}
	l.advancePos()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '+':
		tok = newToken(token.PLUS, string(l.char))
	case '-':
		tok = newToken(token.MINUS, string(l.char))
	case '/':
		tok = newToken(token.SLASH, string(l.char))
	case '*':
		tok = newToken(token.ASTERISK, string(l.char))
	case '!':
		if l.nextTokenIs('=') {
			ch := l.char
			l.advancePos()
			tok = newToken(token.NOT_EQ, string(ch)+string(l.char))
		} else {
			tok = newToken(token.BANG, string(l.char))
		}
	case '=':
		if l.nextTokenIs('=') {
			ch := l.char
			l.advancePos()
			tok = newToken(token.EQ, string(ch)+string(l.char))
		} else {
			tok = newToken(token.ASSIGN, string(l.char))
		}
	case '<':
		if l.nextTokenIs('=') {
			ch := l.char
			l.advancePos()
			tok = newToken(token.LTE, string(ch)+string(l.char))
		} else {
			tok = newToken(token.LT, string(l.char))
		}
	case '>':
		if l.nextTokenIs('=') {
			ch := l.char
			l.advancePos()
			tok = newToken(token.GTE, string(ch)+string(l.char))
		} else {
			tok = newToken(token.GT, string(l.char))
		}
	case ',':
		tok = newToken(token.COMMA, string(l.char))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.char))
	case ':':
		tok = newToken(token.COLON, string(l.char))
	case '(':
		tok = newToken(token.LPAREN, string(l.char))
	case ')':
		tok = newToken(token.RPAREN, string(l.char))
	case '[':
		tok = newToken(token.LBRACKET, string(l.char))
	case ']':
		tok = newToken(token.RBRACKET, string(l.char))
	case '{':
		tok = newToken(token.LBRACE, string(l.char))
	case '}':
		tok = newToken(token.RBRACE, string(l.char))
	case '\'':
		// read char value
		c, cType := l.readChar()
		tok.Literal = c
		tok.Type = cType
	case '"':
		// read string value
		s, sType := l.readString()
		tok.Literal = s
		tok.Type = sType
	case 0:
		tok = newToken(token.EOF, "EOF")
	default:
		if isDigit(l.char) {
			n, nType := l.readNumber()
			tok.Literal = n
			tok.Type = nType
		} else if isLetter(l.char) {
			ident := l.readName()
			indentType := token.LookupIdentType(ident)
			tok.Literal = ident
			tok.Type = indentType
		} else {
			tok = newToken(token.ILLEGAL, string(l.char))
		}
	}
	l.advancePos()

	// DEBUG INFO
	if l.debug {
		fmt.Printf("token: %+v\n", tok)
	}

	return tok
}

func (l *Lexer) advancePos() {
	l.curPos = l.nextPos
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.curPos]
		l.nextPos++
	}
}

func (l *Lexer) nextChar() rune {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) nextTokenIs(ch rune) bool {
	if l.nextPos >= len(l.input) {
		return false
	} else {
		return l.input[l.nextPos] == ch
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' || l.char == '_' {
		l.advancePos()
	}
}

func (l *Lexer) readChar() (string, string) {
	var c string
	var cType string

	pos := l.curPos
	closed := false
	for isLetter(l.nextChar()) || l.nextChar() == '\'' {
		l.advancePos()
		if l.char == '\'' {
			closed = true
		}
	}

	if closed {
		s := string(l.input[pos:l.nextPos])
		sLit := strings.ReplaceAll(s, "'", "")
		if len(sLit) == 1 {
			c = sLit
			cType = token.CHAR_VALUE
		} else {
			c = s
			cType = token.ILLEGAL
		}
	} else {
		c = string(l.input[pos:l.nextPos])
		cType = token.ILLEGAL
	}

	return c, cType
}

func (l *Lexer) readString() (string, string) {
	var s string
	var sType string

	pos := l.curPos
	closed := false
	for l.char != 0 {
		l.advancePos()
		if l.char == '"' {
			closed = true
			break
		}
	}
	s = string(l.input[pos:l.nextPos])

	if closed {
		s = strings.ReplaceAll(s, "\"", "")
		sType = token.STRING_VALUE
	} else {
		sType = token.ILLEGAL
	}

	return s, sType

}

func (l *Lexer) readNumber() (string, string) {
	var nType string
	var n string

	pos := l.curPos
	for isDigit(l.nextChar()) || l.nextChar() == '.' {
		l.advancePos()
	}
	n = string(l.input[pos:l.nextPos])

	if strings.Count(n, ".") == 0 {
		nType = token.INT_VALUE
	} else if strings.Count(n, ".") == 1 {
		nType = token.FLOAT_VALUE
	} else {
		nType = token.ILLEGAL
	}

	return n, nType
}

func (l *Lexer) readName() string {
	var name string

	pos := l.curPos
	for isLetter(l.nextChar()) {
		l.advancePos()
	}
	name = string(l.input[pos:l.nextPos])

	return name
}

func newToken(tokenType string, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch rune) bool {
	return ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '_'
}
