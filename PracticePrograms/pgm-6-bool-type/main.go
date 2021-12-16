package main

import "fmt"

var x bool
var y string

func main() {
	y = "Hello"
	fmt.Printf("%T\n%T\n", x, y)
	fmt.Printf("%s\n", y)
	fmt.Println(y)
	x = true
	fmt.Println(x)

}
