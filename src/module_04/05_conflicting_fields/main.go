package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_04/05_conflicting_fields/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Printf("%T\n\n", organization.TwitterHandler("test"))

	// conflicting types
	fmt.Println("Conflicting Types:")
	p.First = "Collin"
	fmt.Println("p.First:", p.First)
	fmt.Println("p.Name.First:", p.Name.First)
	fmt.Println()

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println()

	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.FullName())
}
