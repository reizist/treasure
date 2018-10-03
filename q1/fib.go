package main

import (
	"fmt"
	"os"
	"strconv"
)

var cache = make(map[int]int, 100)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if m, ok := cache[n]; ok {
		return m
	}
	m := fib(n-2) + fib(n-1)
	cache[n] = m
	return m
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fib(arg))
}
