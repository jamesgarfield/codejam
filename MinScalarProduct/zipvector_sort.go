package main

import (
	"sort"
)

type ZipVectorSorter struct {
	ZipVector
	LessFunc	func([]int64, []int64) bool
}

func (s ZipVectorSorter) Less(i, j int) bool {
	return s.LessFunc(s.ZipVector[i], s.ZipVector[j])
}
func (s ZipVector) Len() int {
	return len(s)
}
func (s ZipVector) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ZipVector) Sort(less func([]int64, []int64) bool) {
	sort.Sort(ZipVectorSorter{s, less})
}
