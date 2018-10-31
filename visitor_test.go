package dp

import (
    "testing"
    "fmt"
)

func TestVisitor(*testing.T) {
    fmt.Printf("Test Visitor:\n\n")

    var e *Element
    var v *Visitor

    e = NewXxxElement("1").Element.SetNext(
        NewZzzElement("-1").Element.SetNext(
            NewXxxElement("-2").Element.SetNext(
                &NewZzzElement("2").Element,
            ),
        ),
    )

    v = &NewXxxVisitor().Visitor
    e.Accept(v)

    v = &NewZzzVisitor().Visitor
    e.Accept(v)

    fmt.Printf("\n")
}