package main

import "fmt"

var x int
var y float64
func main(){
	//x:=42
	//y:=42.567
	x=42
	y=42.567
	fmt.Println(x);
	fmt.Println(y);
	fmt.Printf("%T \n",x);
	fmt.Printf("%T \n",y); //printf basically allows us to print formatted data.%t conversion string 
}