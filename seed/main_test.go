package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

const sample string = ``

func Test_Parse(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))

	expect := 0

	cases := parse(in)
	if len(cases) != expect {
		t.Errorf("Incorrect number of cases. Expected: %d, Found: %d\n", expect, len(cases))
	}
}

func Test_Solve(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(sample))
	cases := parse(in)

	expect := []string{}

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
