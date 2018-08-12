package lexer

import "github.com/roylee0704/monkey/token"

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

func (l *lexer) NextToken() token.Token {

	val := l.input[l.pos : l.pos+1]

	if val == token.PLUS {
		l.pos++
		return token.Token{
			Typ:     token.PLUS,
			Literal: val,
		}
	} else if val == token.EQUAL {
		l.pos++
		return token.Token{
			Typ:     token.EQUAL,
			Literal: val,
		}
	} else if val == token.SEMICOLON {
		l.pos++
		return token.Token{
			Typ:     token.SEMICOLON,
			Literal: val,
		}
	} else if val == token.LBRACE {
		l.pos++
		return token.Token{
			Typ:     token.LBRACE,
			Literal: val,
		}
	} else if val == token.RBRACE {
		l.pos++
		return token.Token{
			Typ:     token.RBRACE,
			Literal: val,
		}
	} else if val == token.LPAREN {
		l.pos++
		return token.Token{
			Typ:     token.LPAREN,
			Literal: val,
		}
	} else if val == token.RPAREN {
		l.pos++
		return token.Token{
			Typ:     token.RPAREN,
			Literal: val,
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
