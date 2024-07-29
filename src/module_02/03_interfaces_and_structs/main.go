package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_02/03_interfaces_and_structs/organization"
)

func main() {
	p := organization.Person{
		FirstName: "James",
		LastName:  "Wilson",
	}

	fmt.Println(p.FirstName, p.LastName)
	p.FirstName = "Collin"
	fmt.Println(p.FirstName, p.LastName)

	fmt.Println("ID:", p.ID())
}
