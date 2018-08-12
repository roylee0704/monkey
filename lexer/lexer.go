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
	pos   int
}

func (l *lexer) consume(tkn string) bool {
	n := len(tkn)
	end := l.pos + n
	val := l.input[l.start:end]

	if val != tkn {
		return false
	}

	l.pos = end
	l.start = l.pos
	return true
}

func (l *lexer) NextToken() token.Token {
	if l.consume(token.PLUS) {
		return token.Token{
			Typ:     token.PLUS,
			Literal: token.PLUS,
		}
	} else if l.consume(token.EQUAL) {
		return token.Token{
			Typ:     token.EQUAL,
			Literal: token.EQUAL,
		}
	} else if l.consume(token.SEMICOLON) {
		return token.Token{
			Typ:     token.SEMICOLON,
			Literal: token.SEMICOLON,
		}
	} else if l.consume(token.LBRACE) {
		return token.Token{
			Typ:     token.LBRACE,
			Literal: "{",
		}
	} else if l.consume(token.RBRACE) {
		return token.Token{
			Typ:     token.RBRACE,
			Literal: "}",
		}
	} else if l.consume(token.LPAREN) {
		return token.Token{
			Typ:     token.LPAREN,
			Literal: token.LPAREN,
		}
	} else if l.consume(token.RPAREN) {
		return token.Token{
			Typ:     token.RPAREN,
			Literal: token.RPAREN,
		}
	} else {
		return token.Token{}
	}
}

// New intantiate a lexer with input
func New(input string) Lexer {
	return &lexer{
		input: input,
		start: 0,
		pos:   0,
	}
}
