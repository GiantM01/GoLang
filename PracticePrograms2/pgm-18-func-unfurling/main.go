package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8}
	x := sum(xi...)
	fmt.Println("The total is", x)
}

func sum(x ...int) int {
	fmt.Println("HELLO WORLD")
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("For item in index position", i, ", We are now adding", v, ",To the total which is now", sum)
	}
	fmt.Println("The total is  ", sum)
	return sum

}
