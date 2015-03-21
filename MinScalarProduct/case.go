package main

import (
	"strconv"
)

//go:generate goast write impl goast.net/x/iter
//go:generate goast write impl goast.net/x/sort

type Vector []int64

func (v Vector) Product() int64 {
	return v.Fold(1, func(a, b int64) int64 { return a * b })
}

func (v Vector) Sum() int64 {
	return v.Fold(0, func(a, b int64) int64 { return a + b })
}

func (v Vector) ScalarProduct(vex ...Vector) int64 {
	var zipped ZipVector = v.Zip(vex...)
	pVec := zipped.EachProduct()
	return pVec.Sum()
}

type ZipVector []Vector

func (z ZipVector) EachProduct() Vector {
	return z.Fold([]int64{}, func(fold, vec []int64) []int64 {
		var v Vector = vec
		return append(fold, v.Product())
	})
}

type Case struct {
	CaseNum  int
	Size     int
	V1       Vector
	V2       Vector
	Solution string
}

func solveCase(c *Case) *Case {

	asc := func(a, b int64) bool { return a < b }
	desc := func(a, b int64) bool { return b < a }
	c.V1.Sort(asc)
	c.V2.Sort(desc)

	c.Solution = strconv.FormatInt(c.V1.ScalarProduct(c.V2), 10)
	return c
}
