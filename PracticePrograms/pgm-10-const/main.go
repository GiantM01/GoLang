package main

import "fmt"

const(
	a int  = 42
	b string = "Mritunjay Upadhyay"
	c bool = true
)
func main(){
    fmt.Println("Value of a is ", a)
		fmt.Println("Value of b is ", b)
		fmt.Println("Value of c is ", c)
		fmt.Printf("%T\n%T\n%T",a,b,c)
}