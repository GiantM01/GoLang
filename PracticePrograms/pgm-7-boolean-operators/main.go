package main

import "fmt"

func main() {
	x := 42
	y := 6
	x = y
	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
}

// if x > y
