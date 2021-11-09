package main

var a = 0

func add(b int) {
	a += b
}

func defer_test() {
	var b = 2
	defer add(b)
	b = 3
}
