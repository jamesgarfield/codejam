package main

import (
	"bufio"
	"strings"
	"testing"
)

const sample string = `3
100
3
5 75 5
200
7
150 24 79 50 88 345 3
8
8
2 1 9 4 4 56 90 3`

func Test_Parse(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)
	if len(cases) != 3 {
		t.Error("Incorrect number of cases. Expected: 3, Found: ", string(len(cases)))
	}
}
