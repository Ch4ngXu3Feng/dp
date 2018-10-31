package dp

import (
    "testing"
    "fmt"
)

func TestComposite(*testing.T) {
    fmt.Printf("Test Composite:\n\n")

    child := NewComposite("child")
    child.Add(&NewLeaf("Child Left").Component)
    child.Add(&NewLeaf("Child Right").Component)

    parent := NewComposite("parent")
    parent.Add(&NewLeaf("Parent Left").Component)
    parent.Add(&NewLeaf("Parent Right").Component)
    parent.Add(&child.Component)

    root := NewComposite("root")
    root.Add(&NewLeaf("Root Left").Component)
    root.Add(&NewLeaf("Root Right").Component)
    root.Add(&parent.Component)

    root.Operation()

    fmt.Printf("\n")
}