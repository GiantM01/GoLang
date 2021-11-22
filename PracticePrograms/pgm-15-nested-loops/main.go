package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("The outer loop of %d\t\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop of %d\t\t\t\t\t\n ", j)
		}
	}
}
