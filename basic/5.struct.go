package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	// cara 1

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	// cara 2
	u2 := user{ID: 1, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u2)

	// cara 2 pake multiple line
	u3 := user{ID: 1,
		FirstName: "Arthur",
		LastName:  "Dent"}
	fmt.Println(u3)
}
