// print name age id salary for 10 employees
package main

import "fmt"

type employee struct {
	name   string
	age    int
	id     int
	salary int
}

func main() {
	emps := []employee{{
		name:   "anurag",
		age:    23,
		salary: 30000,
	}, {
		name:   "mritunjay",
		age:    53,
		salary: 35000,
	},
		{
			name:   "achutyh",
			age:    43,
			salary: 3000000000,
		}}

	// e1 := employee{
	// 	name:   "anurag",
	// 	age:    23,
	// 	salary: 30000,
	// }

	// emps = append(emps, e1)

	for _, k := range emps {
		if k.age > 50 {
			fmt.Printf("%#v \n", k.name)
		}
	}

}
