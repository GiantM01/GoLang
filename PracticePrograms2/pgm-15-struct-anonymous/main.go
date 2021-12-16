package main

import "fmt"

func main() {
	p1 := struct {
		first  string
		second string
		age    int
	}{
		first:  "lionel",
		second: "messi",
		age:    35,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.second, p1.age)
}
