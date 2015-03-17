package main

import (
	"bufio"
	"strconv"
	"strings"
)

type lexer struct {
	Scanner  *bufio.Scanner
	NumCases int
	Cases    []*Case
}

type stateFn func(*lexer) stateFn

func parse(in *bufio.Scanner) []*Case {
	lex := &lexer{Scanner: in}

	for state := parseHeader; state != nil; {
		state = state(lex)
	}

	return lex.Cases
}

func parseHeader(lex *lexer) stateFn {
	if !lex.Scanner.Scan() {
		panic("Invalid input: File Empty")
	}

	lex.NumCases = parseInt(lex.Scanner.Text())

	return parseCase
}

func parseCase(lex *lexer) stateFn {
	if len(lex.Cases) == lex.NumCases {
		return nil
	}

	c := &Case{}

	currentCase := strconv.Itoa(len(lex.Cases) + 1)

	if !lex.Scanner.Scan() {
		panic("Invalid Input: Expected Credit for Case #" + currentCase)
	}
	c.Credit = parseInt(lex.Scanner.Text())

	if !lex.Scanner.Scan() {
		panic("Invalid Input: Expected Items for Case #" + currentCase)
	}
	c.Items = parseInt(lex.Scanner.Text())

	if !lex.Scanner.Scan() {
		panic("Invalid Input: Expected Prices for Case #" + currentCase)
	}
	c.Prices = parsePrices(lex.Scanner.Text())

	lex.Cases = append(lex.Cases, c)

	return parseCase
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic("Invalid Integer: " + s)
	}

	return int(i)
}

func parsePrices(p string) (prices []int) {
	raw := strings.Split(p, " ")
	for _, rp := range raw {
		prices = append(prices, parseInt(rp))
	}
	return
}
