package main

import "fmt"

func main() {
	x := make([]int, 10, 12) // make([]type , length , capacity)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 22)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 23)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 24)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
