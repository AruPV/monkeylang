package lexer

import "monkey/token"

type Lexer struct {
	input 			string
	position 		int // current position in input (cur char's)
	readPosition 	int // current reading position in input (after cur char's)
	ch 				byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhiteSpace()

	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			tok.Literal = "=="
			tok.Type = token.EQ
			lexer.readChar()
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
			lexer.readChar()
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch){
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch){
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string{
	position := lexer.position
	for isDigit(lexer.ch){
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a'<= ch && ch <= 'z' ||'A'<= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}
