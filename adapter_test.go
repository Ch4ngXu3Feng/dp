package dp

import (
    "testing"
    "fmt"
)

func TestAdapter(*testing.T) {
    fmt.Printf("Test Adapter:\n\n")

    var f *File
    var r *Resource
    var ok bool

    adapter := NewAdapter()

    if r, ok = adapter.GetAdapter(r).(*Resource); ok {
        r.Show()
    }

    if f, ok = adapter.GetAdapter(f).(*File); ok {
        f.Show()
    }

    fmt.Printf("\n")
}
