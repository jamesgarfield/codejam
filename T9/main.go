package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type key struct {
	val string
	num int
}

var T9 = map[rune]key{
	' ': {"0", 1},
	'a': {"2", 1},
	'b': {"2", 2},
	'c': {"2", 3},
	'd': {"3", 1},
	'e': {"3", 2},
	'f': {"3", 3},
	'g': {"4", 1},
	'h': {"4", 2},
	'i': {"4", 3},
	'j': {"5", 1},
	'k': {"5", 2},
	'l': {"5", 3},
	'm': {"6", 1},
	'n': {"6", 2},
	'o': {"6", 3},
	'p': {"7", 1},
	'q': {"7", 2},
	'r': {"7", 3},
	's': {"7", 4},
	't': {"8", 1},
	'u': {"8", 2},
	'v': {"8", 3},
	'w': {"9", 1},
	'x': {"9", 2},
	'y': {"9", 3},
	'z': {"9", 4},
}

type Case struct {
	CaseNum  int
	Words    string
	Solution string
}

//go:generate goast write impl github.com/jamesgarfield/sliceops
type Cases []*Case

//go:generate goast write impl pipeline.go
type CaseChan chan *Case

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	in := bufio.NewScanner(file)
	ans := solve(parse(in))
	for _, a := range ans {
		fmt.Println(a)
	}
}

func solve(cases []*Case) (result []string) {
	workers := runtime.NumCPU()
	done := make(chan bool)

	var pipe CaseChan = make(chan *Case)

	go func() {
		for _, c := range cases {
			pipe <- c
		}
	}()

	var solutions Cases = pipe.Fan(done, workers, solveCase).Collect(done, len(cases))

	solutions.Sort(func(a, b *Case) bool { return a.CaseNum < b.CaseNum })

	for _, s := range solutions {
		result = append(result, fmtAns(s))
	}

	return
}

func fmtAns(c *Case) string {
	return fmt.Sprintf("Case #%d: %s", c.CaseNum, c.Solution)
}

func solveCase(c *Case) *Case {
	var lastKey string
	sol := []string{}
	for _, r := range c.Words {
		key := T9[r]
		if lastKey == key.val {
			sol = append(sol, " ")
		}
		lastKey = key.val
		sol = append(sol, strings.Repeat(key.val, key.num))
	}
	c.Solution = strings.Join(sol, "")
	return c
}
