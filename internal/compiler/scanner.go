package compiler

type TokenType int

const (
	BraceClose = iota
	BraceOpen
	Decl
	Eof
	Number
	ParenClose
	ParenOpen
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func Scan(s string) ([]Token, error) {
	return []Token{}, nil
}
