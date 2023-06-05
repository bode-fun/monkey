package lexer_test

import (
	"testing"

	"git.bode.fun/monkey/lexer"
	"git.bode.fun/monkey/token"
	"github.com/matryer/is"
)

func TestNextToken(t *testing.T) {
	is := is.New(t)

	input := `let five = 5;
let ten = 10;

let add = fn(x,y) {
  x + y;
};

let result = add(five, ten);
`

	testTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
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
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := lexer.New(input)

	for _, testToken := range testTokens {
		tok := lexer.NextToken()

		is.Equal(tok.Type, testToken.expectedType)

		is.Equal(tok.Literal, testToken.expectedLiteral)
	}
}
