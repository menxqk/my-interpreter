package lexer

import "github.com/menxqk/my-interpreter/token"

type Lexer struct {
	input []rune

	curPost int
	nextPos int
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	return l
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{Type: token.EOF, Literal: ""}

	return tok
}
