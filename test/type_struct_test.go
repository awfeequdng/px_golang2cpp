package main

type test_type struct {
	a, b int
}

func (tt *test_type) get_a() int {
	return tt.a
}

func (tt *test_type) get_b() int {
	return tt.b
}

type test_cc struct {

	cc string
	dd string
	ee int
}

