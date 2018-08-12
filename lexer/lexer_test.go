package lexer

import (
	"testing"

	"github.com/roylee0704/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = func(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		WantTokenType    token.Type
		WantTokenLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.EQUAL, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.EQUAL, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.EQUAL, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.EQUAL, "="},
		{token.FUNCTION, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, test := range tests {

		tok := l.NextToken()
		if tok.Typ != test.WantTokenType {
			t.Errorf("invalid token type: want: '%v', got: '%v'", test.WantTokenType, tok.Typ)
		}
		if tok.Literal != test.WantTokenLiteral {
			t.Errorf("invalid token literal: want: '%v', got: '%v'", test.WantTokenLiteral, tok.Literal)
		}
	}

	//fmt.Println(tests, input)
}
