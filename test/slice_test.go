package main

func slice_test() int {
	var s1 = []int{1, 2, 34, 5, 6}
	var s2 = s1[:3]
	var ret int = 0
	for _, s := range s2 {
		ret += s
	}
	return ret
}
