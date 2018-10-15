package dp

import (
    "testing"
    "fmt"
)

func TestChain(*testing.T) {
    fmt.Printf("\nTest Chain:\n\n")

    var handler IHandler
    handler = NewXxxHandler(NewZzzHandler(nil))
    handler.Handle()
}
