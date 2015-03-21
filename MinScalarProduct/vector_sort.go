package main

import (
	"sort"
)

type VectorSorter struct {
	Vector
	LessFunc	func(int64, int64) bool
}

func (s VectorSorter) Less(i, j int) bool {
	return s.LessFunc(s.Vector[i], s.Vector[j])
}
func (s Vector) Len() int {
	return len(s)
}
func (s Vector) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Vector) Sort(less func(int64, int64) bool) {
	sort.Sort(VectorSorter{s, less})
}
