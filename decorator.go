package dp

import (
    "fmt"
)

type IVisualComponent interface {
    Draw(text string)
}

type VisualComponent struct {
    IVisualComponent
}

func (c *VisualComponent) Init() {
    fmt.Printf("VisualComponent Init\n")
}

func (c *VisualComponent) Draw(text string) {
    fmt.Printf("VisualComponent Draw %s\n", text)
}

func NewVisualComponent() *VisualComponent {
    c := &VisualComponent{}
    c.Init()
    return c
}

type Decorator struct {
    IVisualComponent
    component IVisualComponent
}

func (d *Decorator) Init() {
    fmt.Printf("Decorator Init\n")
    d.IVisualComponent = d
}

func (d *Decorator) Draw(text string) {
    d.component.Draw(text)
}

func NewDecorator(component IVisualComponent) *Decorator {
    d := &Decorator{
        component: component,
    }
    d.Init()
    return d
}

type XxxDecorator struct {
    Decorator
}

func (d *XxxDecorator) Init() {
    d.Decorator.Init()
    d.IVisualComponent = d
    fmt.Printf("XxxDecorator Init\n")
}

func (d *XxxDecorator) Draw(text string) {
    d.Decorator.Draw(text)
    fmt.Printf("XxxDecorator Draw %s\n", text)
}

func NewXxxDecorator(component IVisualComponent) *XxxDecorator {
    d := &XxxDecorator{
        Decorator: Decorator{
            component: component,
        },
    }
    d.Init()
    return d
}

type ZzzDecorator struct {
    Decorator
}

func (d *ZzzDecorator) Init() {
    d.Decorator.Init()
    d.IVisualComponent = d
    fmt.Printf("ZzzDecorator Init\n")
}

func (d *ZzzDecorator) Draw(text string) {
    d.Decorator.Draw(text)
    fmt.Printf("ZzzDecorator Draw %s\n", text)
}

func NewZzzDecorator(component IVisualComponent) *ZzzDecorator {
    d := &ZzzDecorator{
        Decorator: Decorator{
            component: component,
        },
    }
    d.Init()
    return d
}
