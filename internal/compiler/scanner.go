package compiler

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type TokenType string

const (
	BraceClose  = "BraceClose"
	BraceOpen   = "BraceOpen"
	Declaration = "Declaration"
	Function    = "Function"
	Identifier  = "Identifier"
	EndOfFile   = "EndOfFile"
	Number      = "Number"
	ParenClose  = "ParenClose"
	ParenOpen   = "ParenOpen"
	Plus        = "Plus"
	Print       = "Print"
)

var keywords = map[string]TokenType{
	"fn":    Function,
	"print": Print,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func Scan(input []rune) ([]Token, error) {
	s := scanner{
		input:  input,
		curr:   0,
		tokens: []Token{},
		line:   0,
	}
	return s.scan()
}

type scanner struct {
	input  []rune
	curr   int
	tokens []Token
	line   int
}

func (s *scanner) scan() ([]Token, error) {
	for !s.isEnd() {
		err := s.consumeNext()
		if err != nil {
			return s.tokens, err
		}
	}
	return s.tokens, nil
}

func (s *scanner) isEnd() bool {
	return s.curr >= len(s.input)
}

func (s *scanner) consumeNext() error {
	r := s.input[s.curr]
	if r == '(' {
		newToken := Token{ParenOpen, "(", nil, s.line}
		s.tokens = append(s.tokens, newToken)
		s.curr = s.curr + 1
	} else if r == ')' {
		newToken := Token{ParenClose, ")", nil, s.line}
		s.tokens = append(s.tokens, newToken)
		s.curr = s.curr + 1
	} else if r == '{' {
		newToken := Token{BraceOpen, "{", nil, s.line}
		s.tokens = append(s.tokens, newToken)
		s.curr = s.curr + 1
	} else if r == '}' {
		newToken := Token{BraceClose, "}", nil, s.line}
		s.tokens = append(s.tokens, newToken)
		s.curr = s.curr + 1
	} else if r == '+' {
		newToken := Token{Plus, "+", nil, s.line}
		s.tokens = append(s.tokens, newToken)
		s.curr = s.curr + 1
	} else if r == ':' {
		s.curr = s.curr + 1
		if !s.isEnd() && s.input[s.curr] == '=' {
			s.curr = s.curr + 1
			newToken := Token{Declaration, ":=", nil, s.line}
			s.tokens = append(s.tokens, newToken)
		} else {
			return errors.New(": must be followed by a =")
		}
	} else if unicode.IsLetter(r) {
		start := s.curr
		for !s.isEnd() && unicode.IsLetter(s.input[s.curr]) {
			s.curr = s.curr + 1
		}
		lexeme := string(s.input[start:s.curr])
		tokenType, exists := keywords[lexeme]
		if !exists {
			tokenType = Identifier
		}
		newToken := Token{tokenType, lexeme, nil, s.line}
		s.tokens = append(s.tokens, newToken)
	} else if unicode.IsNumber(r) {
		start := s.curr
		for !s.isEnd() && unicode.IsNumber(s.input[s.curr]) {
			s.curr = s.curr + 1
		}
		lexeme := string(s.input[start:s.curr])
		literal, err := strconv.Atoi(lexeme)
		if err != nil {
			return fmt.Errorf("Failed to convert '%s' to int: %v", lexeme, err)
		}
		newToken := Token{Number, lexeme, literal, s.line}
		s.tokens = append(s.tokens, newToken)
	} else if r == '\n' {
		s.line = s.line + 1
		s.curr = s.curr + 1
	} else if unicode.IsSpace(r) {
		s.curr = s.curr + 1
	} else {
		return fmt.Errorf("Unrecognized character: %c", r)
	}
	return nil
}
