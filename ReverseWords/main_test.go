package main

import (
	"bufio"
	"strings"
	"testing"
)

const sample string = `3
this is a test
foobar
all your base`

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

	expect := []string{
		"Case #1: test a is this",
		"Case #2: foobar",
		"Case #3: base your all",
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
