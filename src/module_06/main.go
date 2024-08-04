package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_06/organization"
)

func main() {
	p := organization.NewPerson(
		"James", "Wilson",
		organization.NewEuropeanUnionIdentifier("My Real ID", "Germany"),
	)

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Printf("%T\n\n", organization.TwitterHandler("test"))

	fmt.Println(p.ID())
	fmt.Println(p.Country())
}
