package main

import (
	"fmt"

	"github.com/koddlers/creating-custom-data-types-go.git/src/module_02/03_interfaces_and_structs/organization"
)

func main() {
	var p organization.Identifiable = organization.Person{}
	fmt.Println(p.ID())
}
