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

    r, ok = adapter.GetAdapter(r).(*Resource)
    if ok {
        r.Show()
    }

    f, ok = adapter.GetAdapter(f).(*File)
    if ok {
        f.Show()
    }

    fmt.Printf("\n")
}
