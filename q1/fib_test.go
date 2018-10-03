package main

import "testing"

func TestFib(t *testing.T) {
	answers := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, v := range answers {
		actual := fib(i)
		if actual != v {
			t.Errorf("got %v\nwant: %v", actual, v)
		}
	}
}
