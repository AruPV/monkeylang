package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T){
	input := `let five = 5;
	let ten = 10;

	let add = fn(x,y) {
		let math = 4 * 3 - 2 + 0/1;
		!num;
		let isit = 3 < 5 > 3;
		if true {
			return 5 == 5;
		} else if false {
			return 5 != 5;
		};
	};
	`

	tests := []struct {
		expectedType 	token.TokenType
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
		{token.LET, "let"},
		{token.IDENT, "math"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.ASTERISK, "*"},
		{token.INT, "3"},
		{token.MINUS, "-"},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "0"},
		{token.SLASH, "/"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.IDENT, "num"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "isit"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.LT, "<"},
		{token.INT, "5"},
		{token.GT, ">"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "5"},
		{token.EQ, "=="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.FALSE, "false"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "5"},
		{token.NOT_EQ, "!="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		println("correct read for type:%q, literal:%q", tt.expectedType, tt.expectedLiteral)
	}
}
