package token

// Type is a token identifier
type Type string

const (
	// LET is a keyword
	LET = "LET"
	// FUNCTION is a keyword
	FUNCTION = "FUNC"

	// IDENT is an identifier
	IDENT = "IDENT"
	// INT is a literal of integer
	INT = "INT"

	// SEMICOLON is a delimeter
	SEMICOLON = ";"
	// LBRACE is a delimeter
	LBRACE = "{"
	// RBRACE is a delimeter
	RBRACE = "}"
	// LPAREN is a delimeter
	LPAREN = "("
	// RPAREN is a delimeter
	RPAREN = ")"

	// PLUS is an arithmetic operator
	PLUS = "PLUS"

	// EQUAL is an arithmetic operator
	EQUAL = "EQUAL"
)