package lexer

import (
	"bufio"
	"fmt"
	"io"
)

var curChar rune
var r *bufio.Reader

func charNext() Token {
	var err error
	if err == io.EOF {
		return TokenEof
	}
	return tokenGet()
}

func tokenGet() Token {
	switch curChar {

	}
	if curChar == '\n' {
		return TokenNewline
	}

	if curChar >= 'a' && curChar <= 'z' {
		word := []rune{curChar}
		for peek() >= 'a' && peek() <= 'z' {
			goNextChar()
			word = append(word, curChar)
		}
		if keyword, ok := checkKeyword(string(word)); ok {
			return keyword
		}
		return TokenIdent
	}
	if curChar == '"' {
		for peek() != '"' {
			goNextChar()
		}
		goNextChar()
		return TokenString
	}
	if curChar >= '0' && curChar <= '9' {
		for peek() >= '0' && peek() <= '9' {
			goNextChar()
		}
		return TokenNumber
	}

	if curChar == '(' {
		return TokenLparen
	}
	if curChar == ')' {
		return TokenLparen
	}

	if curChar == '+' {
		return TokenPlus
	}
	if curChar == '-' {
		return TokenMinus
	}
	if curChar == '*' {
		return TokenTimes
	}
	if curChar == '/' {
		return TokenSlash
	}
	if curChar == '=' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenEql
		default:
			return TokenBecomes
		}
	}
	if curChar == '<' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenLeq
		default:
			return TokenLss
		}
	}
	if curChar == '>' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenGeq
		default:
			return TokenGtr
		}
	}
	panic(fmt.Errorf("invalid symbol: %c", curChar))
}

func checkKeyword(w string) (Token, bool) {
	tok, ok := keywordsMap[w]
	if ok == true {
		return tok, ok
	}
	return -1, false
}

func peek() rune {
	c, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0
		} else {
			panic(err)
		}
	}
	if err = r.UnreadRune(); err != nil {
		panic(err)
	}
	return c
}

func goNextChar() bool {
	c, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return false
		} else {
			panic(err)
		}
	}
	curChar = c
	return true
}

func Lexer(rr *bufio.Reader) []Token {
	r = rr
	tokens := make([]Token, 0)
	for {
		if goNextChar() == false {
			break
		}
		if curChar == ' ' || curChar == '\r' || curChar == '\t' {
			continue
		}
		token := tokenGet()
		tokens = append(tokens, token)
	}
	tokens = append(tokens, TokenEof)
	return tokens
}
