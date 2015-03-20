package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

const sample string = `2
3
1 3 -5
-2 4 1
5
1 2 3 4 5
1 0 1 0 1`

func Test_Parse(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)
	if len(cases) != 2 {
		t.Error("Incorrect number of cases. Expected: 2, Found: ", strconv.Itoa(len(cases)))
	}
}

func Test_Solve(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)

	expect := []string{
		"Case #1: -25",
		"Case #2: 6",
	}

	ans := solve(cases)

	if len(ans) != len(expect) {
		t.Errorf("Incorrect Number of Answers. Expected %d, Found %d", len(expect), len(ans))
		return
	}

	for i, e := range expect {
		if e != ans[i] {
			t.Errorf("Failed Case #%d. Expected %s, Found %s\n", i+1, e, ans[i])
		}
	}
}
