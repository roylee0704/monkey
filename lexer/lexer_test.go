package lexer

import (
	"testing"

	"github.com/roylee0704/monkey/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		WantTokenType    token.Type
		WantTokenLiteral string
	}{
		{token.PLUS, "+"},
		{token.EQUAL, "="},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}

	input := `+=;(){}`

	l := New(input)

	for _, test := range tests {

		itm := l.NextToken()
		if itm.Typ != test.WantTokenType {
			t.Errorf("invalid token type: want: '%v', got: '%v'", test.WantTokenType, itm.Typ)
		}
		if itm.Literal != test.WantTokenLiteral {
			t.Errorf("invalid token literal: want: '%v', got: '%v'", test.WantTokenLiteral, itm.Literal)
		}
	}

	//fmt.Println(tests, input)
}
