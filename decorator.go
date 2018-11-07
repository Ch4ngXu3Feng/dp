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
    d.IVisualComponent = d
    fmt.Printf("Decorator Init\n")
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
    fmt.Printf("XxxDecorator Init\n")
    d.Decorator.Init()
    d.IVisualComponent = d
}

func (d *XxxDecorator) Draw(text string) {
    fmt.Printf("XxxDecorator Draw %s\n", text)
    d.Decorator.Draw(text)
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
    fmt.Printf("ZzzDecorator Init\n")
    d.Decorator.Init()
    d.IVisualComponent = d
}

func (d *ZzzDecorator) Draw(text string) {
    fmt.Printf("ZzzDecorator Draw %s\n", text)
    d.Decorator.Draw(text)
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
