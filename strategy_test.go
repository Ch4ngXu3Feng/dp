package dp

import (
    "testing"
    "fmt"
)

func TestStrategy(*testing.T) {
    fmt.Printf("Test Strategy:\n\n")

    sc := NewStrategyContext()

    sc.SetStrategy(0)
    sc.Do()

    sc.SetStrategy(1)
    sc.Do()

    fmt.Printf("\n")
}
