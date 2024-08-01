package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_04/03_embedding_interfaces_into_structs/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewSocialSecurityNumber("123-45-6789"))

	err := p.SetTwitterHandler("@jam_wils")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
