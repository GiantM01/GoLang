package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Hello")
}

func bar() {
	fmt.Println("Hola")

}
