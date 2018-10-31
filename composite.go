package dp

import "fmt"

type IComponent interface {
    Operation()
    Add(*Component)
    Remove(*Component)
    GetChildren()
}

type Component struct {
    Name string
    children []*Component
    IComponent
}

func (c *Component) Init() {
    c.IComponent = c
    fmt.Printf("Component Init\n")
}

func (c *Component) Operation() {
    if _, ok := c.IComponent.(*Component); ok {
        panic("Component Operation oops\n")
    }
    c.IComponent.Operation()
}

func (c *Component) Add(component *Component) {
    c.children = append(c.children, component)
}

func (c *Component) Remove(component *Component) {
    panic("Component Remove oops\n")
}

func (c *Component) GetChildren() {
    panic("Component GetChildren oops\n")
}

type Composite struct {
    Component
}

func (c *Composite) Init() {
    c.Component.Init()
    c.IComponent = c
    fmt.Printf("Composite Init\n")
}

func (c *Composite) Operation() {
    fmt.Printf("Composite %s Operation\n", c.Name)
    for _, component := range c.children {
        component.Operation()
    }
}

func NewComposite(name string) *Composite {
    c := &Composite{
        Component: Component{
            Name: name,
        },
    }
    c.Init()
    return c
}

type Leaf struct {
    Component
}

func (l *Leaf) Init() {
    l.Component.Init()
    l.IComponent = l
    fmt.Printf("Leaf Init\n")
}

func (l *Leaf) Operation() {
    fmt.Printf("Leaf %s Operation\n", l.Name)
}

func NewLeaf(name string) *Leaf {
    l := &Leaf{
        Component: Component{
            Name: name,
        },
    }
    l.Init()
    return l
}