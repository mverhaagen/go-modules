package mod

import (
	"fmt"
)

type Module struct{}

// Print prints something
func (m *Module) Print() {
	fmt.Println("I'm Here")
}
