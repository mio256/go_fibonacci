package main

import "fmt"

func main() {
	b := 0
	fmt.Println(b)
	c := 1
	fmt.Println(c)
	for a := 10; a > 1; a-- {
		d := b + c
		b = c
		c = d
		fmt.Println(c)
	}
}
