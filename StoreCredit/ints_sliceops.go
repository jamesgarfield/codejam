package main

import (
	"sort"
)

type IntsSorter struct {
	Ints
	LessFunc	func(int, int) bool
}

func (s IntsSorter) Less(i, j int) bool {
	return s.LessFunc(s.Ints[i], s.Ints[j])
}
func (s Ints) Len() int {
	return len(s)
}
func (s Ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Ints) Sort(less func(int, int) bool) {
	sort.Sort(IntsSorter{s, less})
}
func (s Ints) All(fn func(int) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s Ints) Any(fn func(int) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
func (s Ints) Count(fn func(int) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}
func (s Ints) Each(fn func(int)) {
	for _, v := range s {
		fn(v)
	}
}
func (s Ints) First(fn func(int) bool) (match int, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}
func (s Ints) IndexOf(val int) (index int, found bool) {
	for i, v := range s {
		if val == v {
			index = i
			found = true
			break
		}
	}
	return
}
func (s Ints) Where(fn func(int) bool) (result Ints) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s *Ints) Extract(fn func(int) bool) (removed Ints) {
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
