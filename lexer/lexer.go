package lexer

import "waiig/token"

type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	currChar     byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.skipChar()
	return l
}

func (l *Lexer) skipChar() {
	if l.nextPosition >= len(l.input) {
		l.currChar = 0
	} else {
		l.currChar = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.currChar == ' ' || l.currChar == '\t' || l.currChar == '\n' || l.currChar == '\r' {
		l.skipChar()
	}
}
func (l *Lexer) readChar() byte {
	char := l.currChar
	l.skipChar()
	return char
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'B') || char == '_'
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.currPosition
	for isLetter(l.currChar) {
		l.skipChar()
	}
	return l.input[startPosition:l.currPosition]
}

func newToken(typ token.TokenType, char byte) token.Token {
	return token.Token{Type: typ, Literal: string(char)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currChar {
	// Single character
	case '=':
		tok = newToken(token.ASSIGN, l.readChar())
	case ';':
		tok = newToken(token.SEMICOLON, l.readChar())
	case '(':
		tok = newToken(token.LPAREN, l.readChar())
	case ')':
		tok = newToken(token.RPAREN, l.readChar())
	case ',':
		tok = newToken(token.COMMA, l.readChar())
	case '+':
		tok = newToken(token.PLUS, l.readChar())
	case '{':
		tok = newToken(token.LBRACE, l.readChar())
	case '}':
		tok = newToken(token.RBRACE, l.readChar())
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		l.skipChar()

	default:
		if isLetter(l.currChar) { // Identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.IdentType(tok.Literal)
		} else { // Illegal
			tok = newToken(token.ILLEGAL, l.currChar)
		}
	}

	return tok
}
