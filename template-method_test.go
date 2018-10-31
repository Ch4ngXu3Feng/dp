package dp

import (
    "testing"
    "fmt"
)

func TestTemplateMethod(*testing.T) {
    fmt.Printf("Test Template-Method:\n\n")

    ct := NewConcreteTemplate()
    ct.Operation()

    fmt.Printf("\n")
}
