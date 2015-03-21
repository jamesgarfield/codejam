package main

func (s ZipVector) All(fn func([]int64) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s ZipVector) Any(fn func([]int64) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
func (s ZipVector) Count(fn func([]int64) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}
func (s ZipVector) Each(fn func([]int64)) {
	for _, v := range s {
		fn(v)
	}
}
func (s *ZipVector) Extract(fn func([]int64) bool) (removed ZipVector) {
	pos := 0
	kept := *s
	for i := 0; i < len(kept); i++ {
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
func (s ZipVector) First(fn func([]int64) bool) (match []int64, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}
func (s ZipVector) Fold(initial []int64, fn func([]int64, []int64) []int64) []int64 {
	folded := initial
	for _, v := range s {
		folded = fn(folded, v)
	}
	return folded
}
func (s ZipVector) FoldR(initial []int64, fn func([]int64, []int64) []int64) []int64 {
	folded := initial
	for i := len(s) - 1; i >= 0; i-- {
		folded = fn(folded, s[i])
	}
	return folded
}
func (s ZipVector) Where(fn func([]int64) bool) (result ZipVector) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s ZipVector) Zip(in ...ZipVector) (result []ZipVector) {
	minLen := len(s)
	for _, x := range in {
		if len(x) < minLen {
			minLen = len(x)
		}
	}
	for i := 0; i < minLen; i++ {
		row := ZipVector{s[i]}
		for _, x := range in {
			row = append(row, x[i])
		}
		result = append(result, row)
	}
	return
}
