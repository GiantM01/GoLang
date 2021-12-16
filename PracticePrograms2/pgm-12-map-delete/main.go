//to delete a map index delete(m,"key")

package main

import "fmt"

func main() {
	m := map[string]int{
		"Cristiano Ronaldo": 7,
		"Lionel messi ":     10,
	}
	fmt.Println(m)

	// delete(m, "Cristiano Ronaldo")
	// fmt.Println(m)
	for k := range m {
		delete(m, k)
	}
	fmt.Println(m)
}
