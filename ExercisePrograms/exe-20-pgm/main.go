package main

import "fmt"

func main() {
	favSport := "football"
	switch favSport {
	case "cricket":
		fmt.Println("only like watching")

	case "hockey":
		fmt.Println("never played")

	case "football":
		fmt.Println("Love the sport")
	}
}
