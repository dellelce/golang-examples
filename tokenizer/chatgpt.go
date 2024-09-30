package main

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	input string
	pos   int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{input: input}
}

func (t *Tokenizer) nextToken() Token {
	t.skipWhitespace()

	if t.pos >= len(t.input) {
		return Token{Type: EOF}
	}

	switch t.input[t.pos] {
	case '+':
		t.pos++
		return Token{Type: PLUS, Value: "+"}
	case '-':
		t.pos++
		return Token{Type: MINUS, Value: "-"}
	case '*':
		t.pos++
		return Token{Type: MULTIPLY, Value: "*"}
	case '/':
		t.pos++
		return Token{Type: DIVIDE, Value: "/"}
	case '(':
		t.pos++
		return Token{Type: LPAREN, Value: "("}
	case ')':
		t.pos++
		return Token{Type: RPAREN, Value: ")"}
	}

	if unicode.IsDigit(rune(t.input[t.pos])) {
		return t.readNumber()
	}

	t.pos++
	return Token{Type: EOF, Value: string(t.input[t.pos-1])}
}

func (t *Tokenizer) readNumber() Token {
	start := t.pos
	for t.pos < len(t.input) && unicode.IsDigit(rune(t.input[t.pos])) {
		t.pos++
	}
	return Token{Type: NUMBER, Value: t.input[start:t.pos]}
}

func (t *Tokenizer) skipWhitespace() {
	for t.pos < len(t.input) && unicode.IsSpace(rune(t.input[t.pos])) {
		t.pos++
	}
}

func main() {
	input := "3 + 4 * (2 - 1)"
	tokenizer := NewTokenizer(input)

	for {
		token := tokenizer.nextToken()
		fmt.Printf("%v\n", token)
		if token.Type == EOF {
			break
		}
	}
}
