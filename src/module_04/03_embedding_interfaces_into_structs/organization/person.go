package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

// this is a type declaration
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}

func (p *Name) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
}

func (p *Person) ID() string {
	// the `ID()` function exists on both `Identifiable` and `Person`so, `Person` has
	// an `ID()` function that is accesible both via itself and the `Identifiable` interface.
	// Since `ID()` is available via both ways, the one defined on the struct itself
	// (here `Person`) gets precedence.
	return fmt.Sprintf("Persons ID: %s", p.Identifiable.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
