package main

import "fmt"

type person struct {
	fname     string
	lname     string
	favICflav []string
}

func main() {
	p1 := person{
		fname:     "Mritunjay",
		lname:     "Upadhyay",
		favICflav: []string{"chocolate", "vanila", "butterscotch"},
	}
	p2 := person{
		fname:     "Mr",
		lname:     "Devansh",
		favICflav: []string{"strawbeey", "mango", "blackcurrent"},
	}
	fmt.Println(p1.fname)
	fmt.Println(p1.lname)

	for i, v := range p1.favICflav {
		fmt.Println(i, v)
	}
	fmt.Println(p2.fname)
	fmt.Println(p2.lname)

	for i, v := range p1.favICflav {
		fmt.Println(i, v)
	}

}
