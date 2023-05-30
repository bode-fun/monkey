package lexer_test

import (
	"testing"

	"git.bode.fun/monkey/lexer"
	"git.bode.fun/monkey/token"
	"github.com/matryer/is"
)

func TestNextToken(t *testing.T) {
	is := is.New(t)

	input := `=+(){},;`

	testTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
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
