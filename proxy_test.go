package dp

import (
    "testing"
    "fmt"
)

func TestProxy(*testing.T) {
    fmt.Printf("\nTest Proxy:\n\n")

    var subject IObject
    subject = NewObjectProxy()
    subject.Request()
}
