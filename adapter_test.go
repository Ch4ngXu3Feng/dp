package dp

import (
    "testing"
    "fmt"
)

func TestAdapter(*testing.T) {
    fmt.Printf("\nTest Adapter:\n\n")

    adapter := NewAdapter()
    i := (IAdapter)(adapter)
    //t := reflect.TypeOf(i)
    fmt.Printf("Test Adapter addr %x\n", adapter)
    v, ok := i.(* Adapter)
    if ok {
        fmt.Printf("Test Adapter addr %x\n", v)
    }
}
