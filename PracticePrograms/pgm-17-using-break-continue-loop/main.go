package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 { //even
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done")
}
