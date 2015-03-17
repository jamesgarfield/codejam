package main

import (
	"bufio"
	"strings"
	"testing"
)

const sample string = `3
100
3
5 75 25
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

func Test_Solve(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)

	ans := solve(cases)
	if ans[0] != "Case #1: 2 3" {
		t.Error("Failed Case 1, Expected: Case #1: 2 3; Found: ", ans[0])
	}

	if ans[1] != "Case #2: 1 4" {
		t.Error("Failed Case 2, Expected: Case #2: 1 4; Found: ", ans[1])
	}

	if ans[2] != "Case #3: 4 5" {
		t.Error("Failed Case 3, Expected: Case #3: 4 5; Found: ", ans[2])
	}
}
