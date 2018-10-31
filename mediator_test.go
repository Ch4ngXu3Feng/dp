package dp

import (
    "testing"
    "fmt"
)

func TestMediator(*testing.T) {
    fmt.Printf("Test Mediator:\n\n")

    w := NewWidget()
    x := NewXxxWidget()
    z := NewZzzWidget()
    _ = NewMediator(w, x, z)

    w.Changed()
    x.Changed()
    z.Changed()

    fmt.Printf("\n")
}