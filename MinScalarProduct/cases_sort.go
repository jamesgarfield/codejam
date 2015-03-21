package main

import (
	"sort"
)

type CasesSorter struct {
	Cases
	LessFunc	func(*Case, *Case) bool
}

func (s CasesSorter) Less(i, j int) bool {
	return s.LessFunc(s.Cases[i], s.Cases[j])
}
func (s Cases) Len() int {
	return len(s)
}
func (s Cases) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Cases) Sort(less func(*Case, *Case) bool) {
	sort.Sort(CasesSorter{s, less})
}
