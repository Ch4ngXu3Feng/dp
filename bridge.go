package dp

import (
    "fmt"
)

type IGraph interface {
    DrawText()
    DrawLine()
}

type Graph struct {
    IGraph
}

func (g *Graph) Init() {
    fmt.Printf("Graph Init\n")
    g.IGraph = g
}

func (g *Graph) DrawText() {
    if _, ok := g.IGraph.(*Graph); !ok {
        g.IGraph.DrawText()
    }
}

func (g *Graph) DrawLine() {
    if _, ok := g.IGraph.(*Graph); !ok {
        g.IGraph.DrawLine()
    }
}

func NewGraph() *Graph {
    g := &Graph{}
    g.Init()
    return g
}

type XxxGraph struct {
    Graph
}

func (g *XxxGraph) Init() {
    fmt.Printf("XxxGraph Init\n")
    g.Graph.Init()
    g.IGraph = g
}

func (g *XxxGraph) deviceDrawText() {
    fmt.Printf("XxxGraph deviceDrawText\n")
}

func (g *XxxGraph) deviceDrawLine() {
    fmt.Printf("XWindowSystem deviceDrawLine\n")
}

func (g *XxxGraph) DrawText() {
    g.deviceDrawText()
}

func (g *XxxGraph) DrawLine() {
    g.deviceDrawLine()
}

func NewXxxGraph() *XxxGraph {
    x := &XxxGraph{}
    x.Init()
    return x
}

type ZzzGraph struct {
    Graph
}

func (g *ZzzGraph) Init() {
    fmt.Printf("ZzzGraph Init\n")
    g.Graph.Init()
    g.IGraph = g
}

func (g *ZzzGraph) deviceDrawText() {
    fmt.Printf("ZzzGraph deviceDrawText\n")
}

func (g *ZzzGraph) deviceDrawLine() {
    fmt.Printf("ZzzGraph deviceDrawLine\n")
}

func (g *ZzzGraph) DrawText() {
    fmt.Printf("ZzzGraph DrawText\n")
    g.deviceDrawText()
}

func (g *ZzzGraph) DrawLine() {
    fmt.Printf("ZzzGraph DrawLine\n")
    g.deviceDrawLine()
}

func NewZzzGraph() *ZzzGraph {
    g := &ZzzGraph{}
    g.Init()
    return g
}

type Window struct {
    graph IGraph
}

func (w *Window) Init() {
    fmt.Printf("Window Init\n")
}

func (w *Window) DrawText() {
    fmt.Printf("Window DrawText\n")
    w.graph.DrawText()
}

func (w *Window) DrawRect() {
    fmt.Printf("Window DrawRect\n")

    w.graph.DrawLine()
    w.graph.DrawLine()
    w.graph.DrawLine()
    w.graph.DrawLine()
}

func (w *Window) DrawLine() {
    fmt.Printf("Window DrawLine\n")
    w.graph.DrawLine()
}

func NewWindow() *Window {
    w := &Window{}
    w.Init()
    return w
}

type Label struct {
    Window
}

func (l *Label) Init() {
    fmt.Printf("Label Init\n")
    l.Window.Init()
}

func (l *Label) DrawBorder() {
    fmt.Printf("Label DrawBorder\n")

    l.DrawLine()
    l.DrawRect()
    l.DrawText()
}

func NewLabel(graph IGraph) *Label {
    l := &Label{
        Window: Window{
            graph: graph,
        },
    }
    l.Init()
    return l
}

type Button struct {
    Window
}

func (b *Button) Init() {
    fmt.Printf("Button Init\n")
    b.Window.Init()
}

func (b *Button) DrawBorder() {
    fmt.Printf("Button DrawBorder\n")

    b.DrawLine()
    b.DrawRect()
    b.DrawLine()
}

func NewButton(graph IGraph) *Button {
    b := &Button{
        Window: Window{
            graph: graph,
        },
    }
    b.Init()
    return b
}
