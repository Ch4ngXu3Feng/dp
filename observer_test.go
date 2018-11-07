package dp

import (
    "testing"
    "fmt"
)

func TestObserver(*testing.T) {
    fmt.Printf("Test Observer:\n\n")

    s := NewConcreteSubject()

    s.Attach(&NewXxxObserver().Observer)
    s.Attach(&NewZzzObserver().Observer)

    s.SetState(1)
    s.SetState(-1)

    fmt.Printf("\n")
}
