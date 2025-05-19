package main

import (
	"bufio"
	"fmt"
	"strings"

	"small_compiler/lexer"
)

var text = `
if 1 < 2 then
	print "HI"
endif`

func main() {
	r := bufio.NewReader(strings.NewReader(text))
	fmt.Println(lexer.Lexer(r))
}
