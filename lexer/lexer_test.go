package lexer

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

var test_text = `
if 1 < 2 then
	print "HI"
endif`

func TestLexer(t *testing.T) {
	want := []Token{0, 101, 2, 207, 2, 102, 0, 112, 3, 0, 103, -1}
	r := bufio.NewReader(strings.NewReader(test_text))
	tokens := Lexer(r)
	for i := range want {
		if len(tokens) <= i {
			t.Errorf("Got not enough tokens, got only %d right tokens, wanted %d", len(tokens), len(want))
			break
		}
		if want[i] != tokens[i] {
			t.Errorf("Wrong token in place %d. Wanted %s, got %s", i, want[i], tokens[i])
			break
		}
	}
	if len(tokens) > len(want) {
		t.Errorf("Too much tokens. Wanted %d, got %d", len(want), len(tokens))
	}
}

func TestLexerFile1(t *testing.T) {
	want := []Token{101, 2, 209, 2, 102, 0, 112, 3, 0, 103, 0, 0, -1}
	f, err := os.OpenFile("test1", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	tokens := Lexer(r)
	for i := range want {
		if len(tokens) <= i {
			t.Errorf("Got not enough tokens, got only %d right tokens, wanted %d", len(tokens), len(want))
			break
		}
		if want[i] != tokens[i] {
			t.Errorf("Wrong token in place %d. Wanted %s, got %s", i, want[i], tokens[i])
			break
		}
	}
	if len(tokens) > len(want) {
		t.Errorf("Too much tokens. Wanted %d, got %d", len(want), len(tokens))
	}
}

func TestLexerFile2(t *testing.T) {
	want := []Token{112, 3, 0, 111, 1, 0, 0, 107, 1, 110, 2, 0, 107, 1, 110, 2, 0, 104, 1, 209, 2, 105, 0, 112, 1, 0, 107, 1, 110, 1, 201, 1, 0, 107, 1, 110, 1, 0, 107, 1, 110, 1, 0, 107, 1, 110, 1, 202, 2, 0, 106, 0, -1}
	f, err := os.OpenFile("test2", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	tokens := Lexer(r)
	for i := range want {
		if len(tokens) <= i {
			t.Errorf("Got not enough tokens, got only %d right tokens, wanted %d", len(tokens), len(want))
			break
		}
		if want[i] != tokens[i] {
			t.Errorf("Wrong token in place %d. Wanted %s, got %s", i, want[i], tokens[i])
			break
		}
	}
	if len(tokens) > len(want) {
		t.Errorf("Too much tokens. Wanted %d, got %d", len(want), len(tokens))
	}
}
