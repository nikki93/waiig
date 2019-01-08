package lexer

import (
	"testing"
	"waiig/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	expected := []struct {
		typ     token.TokenType
		literal string
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

	l := New(input)

	for i, e := range expected {
		tok := l.NextToken()

		if tok.Type != e.typ {
			t.Fatalf("tests[%d] -- type wrong. expected=%q, got=%q", i, e.typ, tok.Type)
		}

		if tok.Literal != e.literal {
			t.Fatalf("tests[%d] -- literal wrong. expected=%q, got=%q", i, e.literal, tok.Literal)
		}
	}
}
