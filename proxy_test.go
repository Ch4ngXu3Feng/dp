package dp

import (
    "testing"
    "fmt"
)

func TestProxy(*testing.T) {
    fmt.Printf("Test Proxy:\n\n")

    var subject IObject
    subject = NewObjectProxy()
    subject.Request()

    fmt.Printf("\n")
}
