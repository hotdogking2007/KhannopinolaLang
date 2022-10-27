package lexer

import (
	"KhannopinolaLang/token"
	"unicode"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitesapce()
	switch l.ch {
	case rune('='):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case rune('+'):
		tok = newToken(token.PLUS, l.ch)
	case rune('-'):
		tok = newToken(token.MINUS, l.ch)
	case rune('*'):
		tok = newToken(token.ASTERISK, l.ch)
	case rune('!'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case rune('/'):
		if l.peekChar() == rune('/') {
			tok.Type = token.COMMENT
			tok.Literal = l.readString()
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
		tok = newToken(token.SLASH, l.ch)
	case rune('<'):
		tok = newToken(token.LT, l.ch)
	case rune('>'):
		tok = newToken(token.GT, l.ch)
	case rune(';'):
		tok = newToken(token.SEMICOLON, l.ch)
	case rune(','):
		tok = newToken(token.COMMA, l.ch)
	case rune('('):
		tok = newToken(token.LPAREN, l.ch)
	case rune(')'):
		tok = newToken(token.RPAREN, l.ch)
	case rune('{'):
		tok = newToken(token.LBRACE, l.ch)
	case rune('}'):
		tok = newToken(token.RBRACE, l.ch)
	case rune('"'):
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case rune(0):
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return rune(0)
	} else {
		return rune(l.input[l.readPosition])
	}
}
func (l *Lexer) skipWhitesapce() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func (l *Lexer) readIdentifier() string {
	position := l.position
	if unicode.IsLetter(l.ch) {
		l.readChar()
		for unicode.IsLetter(l.ch) || unicode.IsNumber(l.ch) {
			l.readChar()
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == rune('"') || l.ch == rune(0) {
			break
		}
	}

	return string(l.input[position:l.position])
}
func (l *Lexer) readComment() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == rune(0) || l.ch == rune('\n') {
			break
		}
	}

	return string(l.input[position:l.position])
}
func isDigit(ch rune) bool {
	return rune('0') <= ch && ch <= rune('9')
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
