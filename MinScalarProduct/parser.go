package main

import (
	"bufio"
	"strconv"
	"strings"
)

type lexer struct {
	in       *bufio.Scanner
	NumCases int
	Cases    []*Case
}

type stateFn func(*lexer) stateFn

func parse(in *bufio.Scanner) []*Case {
	lex := &lexer{in: in}

	for state := parseHeader; state != nil; {
		state = state(lex)
	}

	return lex.Cases
}

func parseHeader(lex *lexer) stateFn {
	if !lex.in.Scan() {
		panic("Invalid File: Empty")
	}

	lex.NumCases = parseInt(lex.in.Text())

	return parseCase
}

func parseCase(lex *lexer) stateFn {
	if len(lex.Cases) == lex.NumCases {
		return nil
	}

	if !lex.in.Scan() {
		panic("Invalid File: Too Few Cases")
	}

	c := &Case{
		CaseNum: len(lex.Cases) + 1,
		Size:    parseInt(lex.in.Text()),
	}

	if !lex.in.Scan() {
		panic("Invalid File: Missing V1")
	}

	c.V1 = parseVector(lex.in.Text())

	if !lex.in.Scan() {
		panic("Invalid File: Missing V2")
	}

	c.V2 = parseVector(lex.in.Text())

	lex.Cases = append(lex.Cases, c)

	return parseCase
}

func parseVector(s string) (result []int64) {
	raw := strings.Split(s, " ")

	for _, x := range raw {
		result = append(result, parseInt64(x))
	}
	return
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic("Invalid Integer: " + s)
	}

	return int(i)
}

func parseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Invalid Integer: " + s)
	}

	return i
}
