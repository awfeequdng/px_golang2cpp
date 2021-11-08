package main


func condition_for_test() int {
	var a, b = 3, 6
	for a < b {
		a += b
	}
	return a
}

func range_for_test1() int {
	var a = []int{3, 6}
	var b = 0
	for _, i := range a {
		b += i
	}
	return b
}

func range_for_test3() int {
	var a = map[int]int {3: 6, 4: 7,}
	var b = 0
	for _, i := range a {
		b += i
	}
	return b
}

func range_for_test() int {
	var a = []int{3, 6}
	var b = 0
	for i,v := range a {
		b += i
	}
	return b
}

func no_expr_for_test() int {
	var b = 0
	for {
		b += 2
		if b > 10 {
			break
		}
		break
	}
	return b
}