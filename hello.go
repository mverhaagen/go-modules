package mod

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

// Print prints something
func (m *Person) Print() {
	fmt.Println("Person", m)
}
