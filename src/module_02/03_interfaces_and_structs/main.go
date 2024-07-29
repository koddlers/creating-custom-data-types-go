package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_02/03_interfaces_and_structs/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson")

	fmt.Println("ID:", p.ID())
	fmt.Println(p.FullName())
}
