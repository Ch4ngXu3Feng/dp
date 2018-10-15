package dp

import (
    "testing"
    "fmt"
)

func TestAbstractFactory(*testing.T) {
    fmt.Printf("\nTest Abstract Factory:\n\n")

    xf := NewXxxFactory()
    xa := xf.CreateProductA()
    xb := xf.CreateProductB()
    xa.Action()
    xb.Action()

    zf := NewZzzFactory()
    za := zf.CreateProductA()
    zb := zf.CreateProductB()
    za.Action()
    zb.Action()
}
