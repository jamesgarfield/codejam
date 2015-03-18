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
func (s Cases) All(fn func(*Case) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s Cases) Any(fn func(*Case) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
func (s Cases) Count(fn func(*Case) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}
func (s Cases) Each(fn func(*Case)) {
	for _, v := range s {
		fn(v)
	}
}
func (s Cases) First(fn func(*Case) bool) (match *Case, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}
func (s Cases) IndexOf(val *Case) (index int, found bool) {
	for i, v := range s {
		if val == v {
			index = i
			found = true
			break
		}
	}
	return
}
func (s Cases) Where(fn func(*Case) bool) (result Cases) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s *Cases) Extract(fn func(*Case) bool) (removed Cases) {
	pos := 0
	kept := *s
	for i := 0; i < kept.Len(); i++ {
		if fn(kept[i]) {
			removed = append(removed, kept[i])
		} else {
			kept[pos] = kept[i]
			pos++
		}
	}
	kept = kept[:pos:pos]
	*s = kept
	return removed
}
