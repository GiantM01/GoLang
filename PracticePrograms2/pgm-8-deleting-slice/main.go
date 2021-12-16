package main

import "fmt"

func main() {
	x := []int{2, 3, 5, 7, 9, 11}
	fmt.Println(x)
	x = append(x, 13, 15, 16, 19)
	fmt.Println(x)

	y := []int{23, 26, 29, 31}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
