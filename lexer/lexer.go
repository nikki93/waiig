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

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
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

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readNumber() string {
	startPosition := l.currPosition
	for isDigit(l.currChar) {
		l.skipChar()
	}
	return l.input[startPosition:l.currPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currChar {
	// Single character
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.readChar())}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.readChar())}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: string(l.readChar())}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: string(l.readChar())}
	case '<':
		tok = token.Token{Type: token.LT, Literal: string(l.readChar())}
	case '>':
		tok = token.Token{Type: token.GT, Literal: string(l.readChar())}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.readChar())}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.readChar())}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.readChar())}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.readChar())}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.readChar())}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.readChar())}

		// Single or double character
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: string(l.readChar()) + string(l.readChar())}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.readChar())}
		}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: string(l.readChar()) + string(l.readChar())}
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(l.readChar())}
		}

		// EOF
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		l.skipChar()

	default:
		if isLetter(l.currChar) { // Identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.IdentType(tok.Literal)
		} else if isDigit(l.currChar) { // Number
			tok.Literal = l.readNumber()
			tok.Type = token.INT
		} else { // Illegal
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.readChar())}
		}
	}

	return tok
}
