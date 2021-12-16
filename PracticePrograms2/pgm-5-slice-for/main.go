package main

import "fmt"

func main() {
	x := [5]int{2, 4, 6, 8, 42} //COMPOSITE LITERAL
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	// for showing the index and value
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
