package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_03/03_type_definitions/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson")

	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println("ID:", p.ID())
	fmt.Println(p.FullName())
}
