package main

import "fmt"

type batman int
var x batman
var y int 
func main(){
  x = 42
	fmt.Println(x)
	fmt.Printf("%T\n",x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}