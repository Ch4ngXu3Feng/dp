package dp

import "fmt"

type Mediator struct {
    w *Widget
    x *XxxWidget
    z *ZzzWidget
}

func (m *Mediator) Init() {
    m.w.mediator = m
    m.x.mediator = m
    m.z.mediator = m
}

func (m *Mediator) WidgetChanged(w *Widget) {
    switch w {
    case m.w:
        fmt.Printf("Widget Changed:\n")
        m.x.Notify()
        m.z.Notify()
    case &m.x.Widget:
        fmt.Printf("XxxWidget Changed:\n")
        m.z.Notify()
    case &m.z.Widget:
        fmt.Printf("ZzzWidget Changed:\n")
        m.x.Notify()
    }
}

func NewMediator(w *Widget, x *XxxWidget, z *ZzzWidget) *Mediator {
    m := &Mediator{
        w: w,
        x: x,
        z: z,
    }
    m.Init()
    return m
}

type IWidget interface {
    Changed()
}

type Widget struct {
    mediator *Mediator
    IWidget
}

func (w *Widget) Changed() {
    w.mediator.WidgetChanged(w)
}

func NewWidget() *Widget {
    w := &Widget{}
    w.Init()
    return w
}

func (w *Widget) Init() {
    w.IWidget = w
    fmt.Printf("Widget Init\n")
}

type XxxWidget struct {
    Widget
}

func (w *XxxWidget) Notify() {
    fmt.Printf("XxxWidget Notify\n")
}

func NewXxxWidget() *XxxWidget {
    w := &XxxWidget{}
    w.Init()
    return w
}

type ZzzWidget struct {
    Widget
}

func (w *ZzzWidget) Init() {
    w.Widget.Init()
    w.IWidget = w
    fmt.Printf("ZzzWidget Init\n")
}

func (w *ZzzWidget) Notify() {
    fmt.Printf("ZzzWidget Notify\n")
}

func NewZzzWidget() *ZzzWidget {
    w := &ZzzWidget{}
    w.Init()
    return w
}
