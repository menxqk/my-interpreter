package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"

	// Types
	INT_TYPE    = "INT_TYPE"
	FLOAT_TYPE  = "FLOAT_TYPE"
	CHAR_TYPE   = "CHAR_TYPE"
	STRING_TYPE = "STRING_TYPE"

	// Values
	INT_VALUE    = "INT_VALUE"
	FLOAT_VALUE  = "FLOAT_VALUE"
	CHAR_VALUE   = "CHAR_VALUE"
	STRING_VALUE = "STRING_VALUE"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"
	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "!="
	LT       = "<"
	LTE      = "<="
	GT       = ">"
	GTE      = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
)

type Token struct {
	Type    string
	Literal string
}

var dataTypes = map[string]string{
	"int":    INT_TYPE,
	"float":  FLOAT_TYPE,
	"char":   CHAR_TYPE,
	"string": STRING_TYPE,
}

func LookupIdentType(ident string) string {
	if t, ok := dataTypes[ident]; ok {
		return t
	}
	return IDENT
}