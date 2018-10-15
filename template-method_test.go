package dp

import (
    "testing"
    "fmt"
)

func TestTemplateMethod(*testing.T) {
    fmt.Printf("\nTest Template Method:\n\n")

    ct := NewConcreteTemplate()
    ct.Operation()
}
