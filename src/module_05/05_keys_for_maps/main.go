package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_05/05_keys_for_maps/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Printf("%T\n\n", organization.TwitterHandler("test"))

	name1 := Name{First: "", Last: ""}
	// name2 := Name{First: "James", Last: "Wilson"}

	// ssn := organization.NewSocialSecurityNumber("123-45-6789")
	// eu := organization.NewEuropeanUnionIdentifier("12345", "Germany")
	// eu2 := organization.NewEuropeanUnionIdentifier("12345", "Germany")

	portfolio := map[Name][]organization.Person{}
	// using hashable type `Name` as key for our map
	portfolio[name1] = []organization.Person{p}

	if name1 == (Name{}) {
		fmt.Println("We match")
	}
}

type Name struct {
	First string
	Last  string

	// the following will not allow equality check
	// Middle []string
	// Middle func()
	// Middle map[string]string
}

type OtherName struct {
	First string
	Last  string
}
