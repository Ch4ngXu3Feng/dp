package dp

import (
    "testing"
    "fmt"
)

func TestBridge(*testing.T) {
    fmt.Printf("Test Bridge:\n\n")

    var graph IGraph

    graph = NewXxxGraph()
    graph = NewGraph()
    b := NewButton(graph)
    b.DrawBorder()

    graph = NewZzzGraph()
    l := NewLabel(graph)
    l.DrawBorder()

    fmt.Printf("\n")
}