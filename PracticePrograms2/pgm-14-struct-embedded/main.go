package main

import "fmt"

//struct is a sequence of named elements, called "fields", each of which has a name and a type
type player struct {
	fname string // do not put commas after any type it shows a lot of error
	lname string
	age   int
}

type goat struct {
	player
	bdr bool
}

func main() {
	bp := goat{
		player: player{
			fname: "cristiano",
			lname: "ronaldo",
			age:   37,
		},
		bdr: true,
	}

	p2 := player{
		fname: "lionel",
		lname: "messi",
		age:   35,
	}

	fmt.Println(bp)
	fmt.Println(p2)

	fmt.Println(bp.fname, bp.lname, bp.age)
	fmt.Println(p2.fname, p2.lname, p2.age)

}
