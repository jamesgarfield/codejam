package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

const sample string = `4
hi
yes
foo  bar
hello world`

func Test_Parse(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)
	if len(cases) != 4 {
		t.Error("Incorrect number of cases. Expected: 4, Found: ", strconv.Itoa(len(cases)))
	}
}

func Test_Solve(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)

	expect := []string{
		"Case #1: 44 444",
		"Case #2: 999337777",
		"Case #3: 333666 6660 022 2777",
		"Case #4: 4433555 555666096667775553",
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
