package ast

import (
	"KhannopinolaLang/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statement: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
				Type: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "string"},
					Value: "string"
				},
				Kind: 0,
			},
		},
	}

	if program.String() != `let myVar = "anotherVar;"` {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
