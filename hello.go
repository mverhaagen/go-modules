package mod

import (
	"fmt"
)

type Module struct {
	Name string
}

// Print prints something
func (m *Module) Print() {
	fmt.Println("Print", x)
}
