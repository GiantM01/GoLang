package main

import "fmt"

type player struct {
	first string
	last  string
}

type goat struct {
	player
	bdor bool
}

func (g goat) speak() {
	fmt.Println("I am", g.first, g.last)
}

func main() {
	p1 := goat{
		player: player{
			first: "lionel",
			last:  "messi",
		},
		bdor: true,
	}

	fmt.Println(p1)
	p1.speak()
}
