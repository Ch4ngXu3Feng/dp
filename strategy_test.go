package dp

import (
    "testing"
    "fmt"
)

func TestStrategy(*testing.T) {
    fmt.Printf("\nTest Strategy:\n\n")

    sc := NewStrategyContext()

    sc.SetStrategy(0)
    sc.Do()

    sc.SetStrategy(1)
    sc.Do()
}
