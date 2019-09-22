package main

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	NOT      = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ  = "=="
	NEQ = "!="

	//Delims

	COMMA      = ","
	SEMICOLON  = ";"
	LPAREN     = "("
	RPAREN     = ")"
	LBRACKET   = "{"
	RBRACKET   = "}"
	LSQBRACKET = "["
	RSQBRACKET = "]"
	COLON      = ":"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type token_type string

type token struct {
	typ     token_type
	literal string
}

var keywords = map[string]token_type{
	"proc":  FUNCTION,
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
	"ret":   RETURN,
}

func look_up_identifier(ident string) token_type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
