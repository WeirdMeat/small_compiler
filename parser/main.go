package parser

import (
	"fmt"
	"log"

	"github.com/WeirdMeat/small_compiler/lexer"
)

var sym Symbol

func nextsym() {}

func accept(s Symbol) bool {
	if sym == s {
		nextsym()
		return true
	}
	return false
}

func expect(s Symbol) bool {
	if accept(s) {
		return true
	}
	log.Fatal(fmt.Errorf("expect: unexpected symbol"))
	return false
}

func factor() {
	if accept(ident) {

	} else if accept(number) {

	} else if accept(lparen) {
		expression()
		expect(rparen)
	} else {
		log.Fatal(fmt.Errorf("factor: syntax error"))
		nextsym()
	}
}

func term() {
	factor()
	for sym == times || sym == slash {
		nextsym()
		factor()
	}
}

func expression() {
	if sym == plus || sym == minus {
		nextsym()
	}
	term()
	for sym == plus || sym == minus {
		nextsym()
		term()
	}
}

func condition() {
	if accept(oddsym) {
		expression()
	} else {
		expression()
		if sym == eql || sym == neq || sym == lss || sym == leq || sym == gtr || sym == geq {
			nextsym()
			expression()
		} else {
			log.Fatal(fmt.Errorf("condition: invalid operatot"))
			nextsym()
		}
	}
}

func statement() {
	if accept(ident) {
		expect(becomes)
		expression()
	} else if accept(callsym) {
		expect(ident)
	} else if accept(beginsym) {
		for {
			statement()
			if !accept(semicolon) {
				break
			}
		}
		expect(endsym)
	} else if accept(ifsym) {
		condition()
		expect(thensym)
		statement()
	} else if accept(whilesym) {
		condition()
		expect(dosym)
		statement()
	} else {
		log.Fatal(fmt.Errorf("statement: syntax error"))
		nextsym()
	}
}

func block() {
	if accept(constsym) {
		for {
			expect(ident)
			expect(eql)
			expect(number)
			if !accept(comma) {
				break
			}
		}
		expect(semicolon)
	}
	if accept(varsym) {
		for {
			expect(ident)
			if !accept(comma) {
				break
			}
		}
		expect(semicolon)
	}
	for accept(procsym) {
		expect(ident)
		expect(semicolon)
		block()
		expect(semicolon)
	}
	statement()
}

func program() {
	nextsym()
	block()
	expect(period)
}

func main() {
	program()
}
