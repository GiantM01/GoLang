package main

import "fmt"

func main() {
	m := map[string]int{
		"Cristiano Ronaldo": 7,
		"Lionel messi ":     10,
		"Mbappe":            11,
	}
	fmt.Println(m)

	fmt.Println(m["Cristiano Ronaldo"])
	fmt.Println(m["Mbappe"])

	// v, ok := m["Mbappe"]
	// fmt.Println(v)
	// fmt.Println(ok)

	if v, ok := m["Mbappe"]; ok {
		fmt.Println(v)
	}

	m["Neymar"] = 11

	for k, v := range m {
		fmt.Println(k, v)
	}

}
