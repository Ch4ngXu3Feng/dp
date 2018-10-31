package dp

import (
    "testing"
    "fmt"
)

func TestChain(*testing.T) {
    fmt.Printf("Test Chain:\n\n")

    var handler IHandler
    handler = NewXxxHandler(NewZzzHandler(nil))
    handler.Handle()

    fmt.Printf("\n")
}
