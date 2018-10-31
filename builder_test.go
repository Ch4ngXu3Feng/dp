package dp

import (
    "testing"
    "fmt"
)

func TestBuilder(*testing.T) {
    fmt.Printf("Test Builder:\n\n")

    var product *Product
    var builder IBuilder
    d := NewDirector()

    builder = NewXxxBuilder()
    product = d.Create(builder)
    product.Show()

    builder = NewZzzBuilder()
    product = d.Create(builder)
    product.Show()

    fmt.Printf("\n")
}
