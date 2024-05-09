// token/token.go

package token

// Type defines the type of a token. Using a string to distinguish between "integer" and
// "right bracket" more easily.
type Type string

// Token represents a single token, it holds the type of the token, as well as a Literal
// field to store data for later use.
type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // foo, bar, baz, x, y
	INT   = "INT"   // 123456

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NE = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,

	"true":  TRUE,
	"false": FALSE,

	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
