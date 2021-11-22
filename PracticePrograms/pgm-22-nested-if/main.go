package main

import "fmt"

func main() {
	x := 22
	if x == 20 {
		fmt.Println("Our code was 20")
	} else if x == 21 {
		fmt.Println("Our code was 21")
	} else if x == 22 {
		fmt.Println("Our code was 22")
	} else {
		fmt.Println("Our code is neither 20,21,22")
	}
}
