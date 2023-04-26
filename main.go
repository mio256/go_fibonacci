package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	b := 0
	c := 1
	for a := n; a > 1; a-- {
		d := b + c
		b = c
		c = d
	}
	return c
}

func main() {
	fmt.Println(fib(5))
}
