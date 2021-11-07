package main

func condition_for_test() int {
	a, b := 3, 6
	for a < b {
		a += b
	}
	return a
}
