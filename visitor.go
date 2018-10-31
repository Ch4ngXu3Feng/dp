package dp

import "fmt"

type IVisitor interface {
    VisitXxxElement(*Element)
    VisitZzzElement(*Element)
}

type Visitor struct {
    IVisitor
}

func (v *Visitor) Init() {
    v.IVisitor = v
    fmt.Printf("Visitor Init\n")
}

func (v *Visitor) VisitXxxElement(x *Element) {
    if _, ok := v.IVisitor.(*Visitor); ok {
        panic("Visitor VisitXxxElement oops\n")
    }
    v.IVisitor.VisitXxxElement(x)
}

func (v *Visitor) VisitZzzElement(z *Element) {
    if _, ok := v.IVisitor.(*Visitor); ok {
        panic("Visitor VisitZzzElement oops\n")
    }
    v.IVisitor.VisitZzzElement(z)
}

func NewVistor() *Visitor {
    v := &Visitor{}
    v.Init()
    return v
}

type XxxVisitor struct {
    Visitor
}

func (v *XxxVisitor) Init() {
    v.Visitor.Init()
    v.IVisitor = v
    fmt.Printf("XxxVisitor Init\n")
}

func (v *XxxVisitor) VisitXxxElement(e *Element) {
    x, ok := e.IElement.(*XxxElement)
    if ok {
        fmt.Printf("XxxVisitor VisitXxxElement:\n")
        x.Show()
    }
}

func (v *XxxVisitor) VisitZzzElement(e *Element) {
    z, ok := e.IElement.(*ZzzElement)
    if ok {
        fmt.Printf("XxxVisitor VisitZzzElement:\n")
        z.Show()
    }
}

func NewXxxVisitor() *XxxVisitor {
    v := &XxxVisitor{}
    v.Init()
    return v
}

type ZzzVisitor struct {
    Visitor
}

func (v *ZzzVisitor) Init() {
    v.Visitor.Init()
    v.IVisitor = v
    fmt.Printf("ZzzVisitor Init\n")
}

func (v *ZzzVisitor) VisitXxxElement(e *Element) {
    x, ok := e.IElement.(*XxxElement)
    if ok {
        fmt.Printf("ZzzVisitor VisitXxxElement:\n")
        x.Show()
    }
}

func (v *ZzzVisitor) VisitZzzElement(e *Element) {
    z, ok := e.IElement.(*ZzzElement)
    if ok {
        fmt.Printf("ZzzVisitor VisitXxxElement:\n")
        z.Show()
    }
}

func NewZzzVisitor() *ZzzVisitor {
    v := &ZzzVisitor{}
    v.Init()
    return v
}

type IElement interface {
    Accept(*Visitor)
    Show()
}

type Element struct {
    Name string
    Next *Element
    IElement
}

func (e *Element) Init() {
    e.IElement = e
    fmt.Printf("Element Init\n")
}

func (e *Element) Accept(v *Visitor) {
    if _, ok := e.IElement.(*Element); ok {
        panic("Element Accept oops\n")
    }
    e.IElement.Accept(v)
}

func (e *Element) Show() {
    if _, ok := e.IElement.(*Element); ok {
        panic("Element Show oops\n")
    }
    e.IElement.Show()
}

func (e *Element) SetNext(next *Element) *Element {
    e.Next = next
    return e
}

func (e *Element) CallNext(v *Visitor) {
    if e.Next != nil {
        e.Next.Accept(v)
    }
}

func NewElement() *Element {
    e := &Element{}
    e.Init()
    return e
}

type XxxElement struct {
    Element
}

func (e *XxxElement) Init() {
    e.Element.Init()
    e.IElement = e
    fmt.Printf("XxxElement Init\n")
}

func (e *XxxElement) Accept(v *Visitor) {
    v.VisitXxxElement(&e.Element)
    e.CallNext(v)
}

func (e *XxxElement) Show() {
    fmt.Printf("XxxElement is %s\n", e.Name)
}

func NewXxxElement(name string) *XxxElement {
    e := &XxxElement{
        Element: Element{
            Name: name,
        },
    }
    e.Init()
    return e
}

type ZzzElement struct {
    Element
}

func (e *ZzzElement) Init() {
    e.Element.Init()
    e.IElement = e
    fmt.Printf("ZzzElement Init\n")
}

func (e *ZzzElement) Accept(v *Visitor) {
    v.VisitZzzElement(&e.Element)
    e.CallNext(v)
}

func (e *ZzzElement) Show() {
    fmt.Printf("ZzzElement is %s\n", e.Name)
}

func NewZzzElement(name string) *ZzzElement {
    e := &ZzzElement{
        Element: Element{
            Name: name,
        },
    }
    e.Init()
    return e
}
