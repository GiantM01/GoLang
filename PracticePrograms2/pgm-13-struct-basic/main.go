package main

import "fmt"

type player struct {
	fname string
	lname string
	age   int
}

func main() {

	p1 := player{
		fname: "cristiano",
		lname: "ronaldo",
		age:   37,
	}

	p2 := player{
		fname: "lionel",
		lname: "messi",
		age:   35,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname, p2.lname, p2.age)

	boots()

}
