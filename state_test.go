package dp

import (
    "testing"
    "fmt"
)

func TestState(*testing.T) {
    fmt.Printf("Test State:\n\n")

    var context *StateContext

    Xxx := NewXxxState()
    context = NewStateContext(&Xxx.State)
    context.Request()

    Zzz := NewZzzState()
    context = NewStateContext(&Zzz.State)
    context.Request()

    fmt.Printf("\n")
}
