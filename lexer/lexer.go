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
	if l.next >= len(l.input) {
		l.tok = 0
	} else {
		l.tok = l.input[l.next]
	}
	// todo(roy): there is a better way to express this, robpike way.
	// such that it could remove the duplicate `var start int` in every
	// fn: readXXX()
	l.start = l.next
	l.next = l.next + 1
}

// chom numbers
func (l *lexer) readNumber() string {
	var start = l.start
	for isDigit(l.tok) {
		l.readChar()
	}
	return l.input[start:l.start]
}

// chom identifiers (do not support alphanumeric)
func (l *lexer) readIdentifier() string {
	var start = l.start
	for isLetter(l.tok) {
		l.readChar()
	}
	return l.input[start:l.start]
}

func (l *lexer) skipWhitespaces() {
	for isWhitespace(l.tok) {
		l.readChar()
	}
}

func (l *lexer) NextToken() token.Token {
	l.skipWhitespaces()

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
	case 0:
		tok.Typ = token.EOF
		tok.Literal = ""

	// todo: on case of multicharacter: identifier, keyword, integer
	// todo: skipping whitespaces
	default:
		if isDigit(l.tok) {
			tok.Typ = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.tok) {
			val := l.readIdentifier()
			if val == "let" {
				tok.Typ = token.LET
				tok.Literal = val
				return tok
			} else if val == "func" {
				tok.Typ = token.FUNCTION
				tok.Literal = val
				return tok
			} else {
				tok.Typ = token.IDENT
				tok.Literal = val
				return tok
			}

		} else {
			tok.Typ = token.ILLEGAL
			tok.Literal = "line ?"
		}
	}

	l.readChar()
	return tok
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z'
}

// New intantiate a lexer with input
func New(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}
