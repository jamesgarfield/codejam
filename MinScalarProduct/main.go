package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

const MAXVAL = 100001
const MINVAL = -100001

type Case struct {
	CaseNum  int
	Size     int
	V1       Vector
	V2       Vector
	Solution string
}

//go:generate goast write impl github.com/jamesgarfield/sliceops
type Cases []*Case
type Vector []int64

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

	asc := func(a, b int64) bool { return a < b }
	desc := func(a, b int64) bool { return b < a }
	c.V1.Sort(asc)
	c.V2.Sort(desc)

	c.Solution = strconv.FormatInt(scalarProduct(c.V1, c.V2), 10)
	return c
}

func solveCaseA(c *Case) *Case {
	V1min1, V1min2, V1minVal1, _ := mins(c.V1)
	V1max1, V1max2, V1maxVal1, _ := maxs(c.V1)

	V2min1, V2min2, V2minVal1, _ := mins(c.V2)
	V2max1, V2max2, V2maxVal1, _ := maxs(c.V2)

	deltaA := (V1maxVal1 - V2minVal1)
	deltaB := (V2maxVal1 - V1minVal1)

	var minVec []int64
	var max1, max2, min1, min2 int
	if deltaA > deltaB {
		max1, max2 = V1max1, V1max2
		min1, min2 = V2min1, V2min2
		minVec = c.V2
	} else {
		max1, max2 = V2max1, V2max2
		min1, min2 = V1min1, V1min2
		minVec = c.V1
	}

	minVec[max1], minVec[min1] = minVec[min1], minVec[max1]
	minVec[max2], minVec[min2] = minVec[min2], minVec[max2]

	c.Solution = strconv.FormatInt(scalarProduct(c.V1, c.V2), 10)
	return c
}

func scalarProduct(vec1, vec2 []int64) int64 {
	p := int64(0)
	for i, v1 := range vec1 {
		v2 := vec2[i]
		p += (v1 * v2)
	}

	return p
}

func mins(vec []int64) (min1, min2 int, minVal1, minVal2 int64) {
	minVal1 = int64(MAXVAL)
	minVal2 = int64(MAXVAL)

	for i, v := range vec {
		if v >= minVal1 && v >= minVal2 {
			continue
		}
		//Lowest value yet, bump up
		if v < minVal1 {
			minVal1, minVal2 = v, minVal1
			min1, min2 = i, min1
			continue
		}

		//Lower than minval2
		if v < minVal2 {
			minVal2 = v
			min2 = i
		}
	}

	return

}

func maxs(vec []int64) (max1, max2 int, maxVal1, maxVal2 int64) {
	maxVal1 = int64(MINVAL)
	maxVal2 = int64(MINVAL)

	for i, v := range vec {
		if v <= maxVal1 && v <= maxVal2 {
			continue
		}

		//Highest value yet
		if v > maxVal1 {
			maxVal1, maxVal2 = v, maxVal1
			max1, max2 = i, max1
			continue
		}

		if v > maxVal2 {
			maxVal2 = v
			max2 = i
		}
	}

	return
}
