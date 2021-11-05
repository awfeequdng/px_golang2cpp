package if_stmt_test

func add(a, b int) int {
	return a + b
}

func if_stmt_test1() int {
	if v := add(2, add(5, 7)); v > 20 {
		return v;
	} else {
		t := add(3, 7)
		return t
	}
	return 0
}

func two_returns(a int) (string, int) {
	if a > 0 {

		return "a: ", add(a, 10)
	} else {
		return "", a
	}
}

func if_stmt_test2() int {
	if s, v := two_returns(2); v > 0 {
		return v
	}
	return 0
}

func if_stmt_test3() int {
	if s, v := two_returns(3); v > 0 {
		return v
	} else if v == 0 {
		return 0
	} else {
		return -v
	}
}