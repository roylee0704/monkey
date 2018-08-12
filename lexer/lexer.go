package lexer

import (
	"github.com/roylee0704/monkey/token"
)

// Lexer provides an interface that tokenise
// an array of characters
type Lexer interface {
	// NextToken scans and return next recognisable
	// token.
	NextToken() token.Token
}

type lexer struct {
	input string
	start int
	next  int

	tok byte
}

// readChar advances `next` by reading 1 byte to `tok`
// from `input` at a time.
//
// On top of that, byte '0' is assigned to tok if `next`
// hits the limit of input length.
func (l *lexer) readChar() {
	if l.next > len(l.input) {
		l.tok = 0
	} else {
		l.tok = l.input[l.next]
	}

	l.start = l.next
	l.next = l.next + 1
}

func (l *lexer) NextToken() token.Token {
	l.readChar()

	var tok token.Token
	switch l.tok {
	case '+':
		tok = token.New(token.PLUS, l.tok)
	case '=':
		tok = token.New(token.EQUAL, l.tok)
	case ';':
		tok = token.New(token.SEMICOLON, l.tok)
	case '{':
		tok = token.New(token.LBRACE, l.tok)
	case '}':
		tok = token.New(token.RBRACE, l.tok)
	case '(':
		tok = token.New(token.LPAREN, l.tok)
	case ')':
		tok = token.New(token.RPAREN, l.tok)
	case ',':
		tok = token.New(token.COMMA, l.tok)

	// todo: on case of multicharacter: identifier, keyword, integer
	// todo: skipping whitespaces
	case 0:
		tok.Typ = token.EOF
		tok.Literal = ""
	default:
		tok.Typ = token.ILLEGAL
		tok.Literal = "line ?"
	}
	return tok
}

// New intantiate a lexer with input
func New(input string) Lexer {
	return &lexer{input: input}
}
