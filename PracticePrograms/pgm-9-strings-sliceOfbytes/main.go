package main

import "fmt"

func main() {
	s := "Hello, I am Mritunjay Upadhyay"
	fmt.Println(s)
	fmt.Printf("%T \n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Println("at index position %d we have %#x", i, v)
	}
}
