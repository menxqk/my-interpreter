package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

type Token struct {
	Type    string
	Literal string
}
