package dp

import (
    "testing"
    "fmt"
)

func TestFactoryMethod(*testing.T) {
    fmt.Printf("Test Factory-Method:\n\n")

    var c *Creator
    c = &NewXxxCreator().Creator
    c.NewDocument("Xxx")
    c.NewDocument("Zzz")
    c.Show()

    c = &NewZzzCreator().Creator
    c.NewDocument("Xxx")
    c.NewDocument("Zzz")
    c.Show()

    fmt.Printf("\n")
}