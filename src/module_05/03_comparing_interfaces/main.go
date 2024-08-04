package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_05/03_comparing_interfaces/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Printf("%T\n\n", organization.TwitterHandler("test"))

	name1 := Name{First: "James", Last: "Wilson"}
	name2 := Name{First: "James", Last: "Wilson"}
	// name3 := OtherName{First: "James", Last: "Wilson"}

	if name1 == name2 {
		fmt.Println("Names match")
	}

	// ERROR: cannot compare type `Name` with type `OtherName`
	// if name1 == name3 {
	// 	fmt.Println("Names match")
	// }

	ssn := organization.NewSocialSecurityNumber("123-45-6789")
	eu := organization.NewEuropeanUnionIdentifier("12345", "Germany")

	if ssn == eu {
		fmt.Println("Interfaces match!!!")
	} else {
		fmt.Println("Unfortunately, those Interfaces, they don't match!")
	}

	// uncommenting the following will produce error and the program will panic
	// eu2 := organization.NewEuropeanUnionIdentifier("12345", "Germany")
	// if eu == eu2 {
	// 	fmt.Println("Those EU Interfaces, they match. No big surprises there.")
	// }
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
