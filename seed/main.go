package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Case struct {
	CaseNum  int
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

	return c
}
