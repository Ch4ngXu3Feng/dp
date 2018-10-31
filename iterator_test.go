package dp

import (
    "testing"
    "fmt"
)

func TestIterator(*testing.T) {
    fmt.Printf("Test Iterator:\n\n")

    elements := NewElements("Xxx", "Zzz")
    i := &NewElementsIterator(elements).Iterator
    for e := i.First(); e != i.Last(); e = i.Next() {
        fmt.Printf("Elements: %s\n", e.(string))
    }

    fmt.Printf("\n")
}
