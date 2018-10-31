package dp

import (
    "fmt"
    "testing"
)

func TestDecorator(*testing.T) {
    fmt.Printf("Test Decorator:\n\n")

    var vc IVisualComponent
    vc = NewXxxDecorator(NewZzzDecorator(NewVisualComponent()))
    vc.Draw("Test Decorator")

    fmt.Printf("\n")
}
