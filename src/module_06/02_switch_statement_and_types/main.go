package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_06/02_switch_statement_and_types/organization"
)

func main() {
	p := organization.NewPerson(
		"James", "Wilson",
		organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"),
	)

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Printf("%T\n\n", organization.TwitterHandler("test"))
}
