package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"

	// Types
	INT_TYPE    = "INT"
	FLOAT_TYPE  = "FLOAT"
	CHAR_TYPE   = "CHAR"
	STRING_TYPE = "STRING"

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

	// Keywords
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"
)

type Token struct {
	Type    string
	Literal string
}

var keywords = map[string]string{
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"null":   NULL,
	"false":  FALSE,
	"int":    INT_TYPE,
	"float":  FLOAT_TYPE,
	"char":   CHAR_TYPE,
	"string": STRING_TYPE,
}

var dataTypes = map[string]string{
	"int":    INT_TYPE,
	"float":  FLOAT_TYPE,
	"char":   CHAR_TYPE,
	"string": STRING_TYPE,
}

func LookupIdentType(ident string) string {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}

func IsDataType(tokenType string) bool {
	if _, ok := dataTypes[tokenType]; ok {
		return true
	}
	return false
}
