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
func (s Vector) All(fn func(int64) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s Vector) Any(fn func(int64) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
func (s Vector) Count(fn func(int64) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}
func (s Vector) Each(fn func(int64)) {
	for _, v := range s {
		fn(v)
	}
}
func (s Vector) First(fn func(int64) bool) (match int64, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}
func (s Vector) IndexOf(val int64) (index int, found bool) {
	for i, v := range s {
		if val == v {
			index = i
			found = true
			break
		}
	}
	return
}
func (s Vector) Where(fn func(int64) bool) (result Vector) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s *Vector) Extract(fn func(int64) bool) (removed Vector) {
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
