package lexer

import "git.bode.fun/monkey/token"

type Lexer struct {
	input string
	// The position of the head (current char)
	position int
	// Reading position (next char)
	readPosition int
	// Character
	char byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	// Read the initial character
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	default:
		if isLetter(l.char) {
			// Keywords and identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			// Numbers
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// Unknown tokens
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// Advance the pointer until the next character
func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readNumber() string {
	startPosition := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
