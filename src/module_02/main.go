package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_02/03_interfaces_and_structs/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson")

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println("ID:", p.ID())
	fmt.Println(p.FullName())
}
