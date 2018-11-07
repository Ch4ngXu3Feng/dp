package dp

import (
    "testing"
    "fmt"
)

func TestFlyweight(*testing.T) {
    fmt.Printf("Test Flyweight:\n\n")

    var f *Flyweight
    fm := NewFlyweightManager()

    f = fm.GetFlyweight("Xxx")
    f.Operation()

    f = fm.GetFlyweight("Zzz")
    f.Operation()

    fmt.Printf("\n")
}