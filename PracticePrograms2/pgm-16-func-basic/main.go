package main

import "fmt"

func main() {
	foo()
	bar("monica")
}

func foo() {
	fmt.Println("Hello from FOO!")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}
