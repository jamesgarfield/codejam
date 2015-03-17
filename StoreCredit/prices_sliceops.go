package main

import (
	"sort"
)

type PricesSorter struct {
	Prices
	LessFunc	func(*Price, *Price) bool
}

func (s PricesSorter) Less(i, j int) bool {
	return s.LessFunc(s.Prices[i], s.Prices[j])
}
func (s Prices) Len() int {
	return len(s)
}
func (s Prices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Prices) Sort(less func(*Price, *Price) bool) {
	sort.Sort(PricesSorter{s, less})
}
func (s Prices) All(fn func(*Price) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s Prices) Any(fn func(*Price) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
func (s Prices) Count(fn func(*Price) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}
func (s Prices) Each(fn func(*Price)) {
	for _, v := range s {
		fn(v)
	}
}
func (s Prices) First(fn func(*Price) bool) (match *Price, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}
func (s Prices) IndexOf(val *Price) (index int, found bool) {
	for i, v := range s {
		if val == v {
			index = i
			found = true
			break
		}
	}
	return
}
func (s Prices) Where(fn func(*Price) bool) (result Prices) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s *Prices) Extract(fn func(*Price) bool) (removed Prices) {
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
