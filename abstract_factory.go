package dp

import "fmt"

type IAbstractFactory interface {
    CreateProductA() *AbstractProductA
    CreateProductB() *AbstractProductB
}

type AbstractFactory struct {
    IAbstractFactory
}

func (f *AbstractFactory) Init() {
    fmt.Printf("AbstractFactory Init\n")
    f.IAbstractFactory = f
}

func (f *AbstractFactory) CreateProductA() *AbstractProductA {
    if _, ok := f.IAbstractFactory.(*AbstractFactory); ok {
        panic("AbstractFactory CreateProductA oops\n")
    }
    return f.IAbstractFactory.CreateProductA()
}

func (f *AbstractFactory) CreateProductB() *AbstractProductB {
    if _, ok := f.IAbstractFactory.(*AbstractFactory); ok {
        panic("AbstractFactory CreateProductB oops\n")
    }
    return f.IAbstractFactory.CreateProductB()
}

func NewAbstractFactory() *AbstractFactory {
    f := &AbstractFactory{}
    f.Init()
    return f
}

type XxxFactory struct {
    AbstractFactory
}

func (f *XxxFactory) Init() {
    fmt.Printf("XxxFactory Init\n")
    f.AbstractFactory.Init()
    f.IAbstractFactory = f
}

func (f *XxxFactory) CreateProductA() *AbstractProductA {
    fmt.Printf("XxxFactory CreateProductA\n")
    return &NewZzzProductA().AbstractProductA
}

func (f *XxxFactory) CreateProductB() *AbstractProductB {
    fmt.Printf("ZzzFactory CreateProductB\n")
    return &NewXxxProductB().AbstractProductB
}

func NewXxxFactory() *XxxFactory {
    f := &XxxFactory{}
    f.Init()
    return f
}

type ZzzFactory struct {
    AbstractFactory
}

func (f *ZzzFactory) Init() {
    fmt.Printf("ZzzFactory Init\n")
    f.AbstractFactory.Init()
    f.IAbstractFactory = f
}

func (f *ZzzFactory) CreateProductA() *AbstractProductA {
    fmt.Printf("ZzzFactory CreateProductA\n")
    return &NewZzzProductA().AbstractProductA
}

func (f *ZzzFactory) CreateProductB() *AbstractProductB {
    fmt.Printf("ZzzFactory CreateProductB\n")
    return &NewZzzProductB().AbstractProductB
}

func NewZzzFactory() *ZzzFactory {
    f := &ZzzFactory{}
    f.Init()
    return f
}

type IAbstractProductA interface {
    Action()
}

type AbstractProductA struct {
    IAbstractProductA
}

func (a *AbstractProductA) Init() {
    fmt.Printf("AbstractProductA Init\n")
    a.IAbstractProductA = a
}

func (a *AbstractProductA) Action() {
    a.IAbstractProductA.Action()
}

func NewAbstractProductA() *AbstractProductA {
    a := &AbstractProductA{}
    a.Init()
    return a
}

type XxxProductA struct {
    AbstractProductA
}

func (a *XxxProductA) Init() {
    fmt.Printf("XxxProductA Init\n")
    a.AbstractProductA.Init()
    a.IAbstractProductA = a
}

func (a *XxxProductA) Action() {
    fmt.Printf("XxxProductA Action\n")
}

func NewXxxProductA() *XxxProductA {
    a := &XxxProductA{}
    a.Init()
    return a
}

type ZzzProductA struct {
    AbstractProductA
}

func (a *ZzzProductA) Init() {
    fmt.Printf("ZzzProductA Init\n")
    a.AbstractProductA.Init()
    a.IAbstractProductA = a
}

func (a *ZzzProductA) Action() {
    fmt.Printf("ZzzProductA Action\n")
}

func NewZzzProductA() *ZzzProductA {
    a := &ZzzProductA{}
    a.Init()
    return a
}

type IAbstractProductB interface {
    Action()
}

type AbstractProductB struct {
    IAbstractProductB
}

func (b *AbstractProductB) Init() {
    fmt.Printf("AbstractProductB Init\n")
    b.IAbstractProductB = b
}

func (b *AbstractProductB) Action() {
    b.IAbstractProductB.Action()
}

func NewAbstractProductB() *AbstractProductB {
    b := &AbstractProductB{}
    b.Init()
    return b
}

type XxxProductB struct {
    AbstractProductB
}

func (b *XxxProductB) Init() {
    fmt.Printf("XxxProductB Init\n")
    b.AbstractProductB.Init()
    b.IAbstractProductB = b
}

func (b *XxxProductB) Action() {
    fmt.Printf("XxxProductB Action\n")
}

func NewXxxProductB() *XxxProductB {
    b := &XxxProductB{}
    b.Init()
    return b
}

type ZzzProductB struct {
    AbstractProductB
}

func (b *ZzzProductB) Init() {
    fmt.Printf("ZzzProductB Init\n")
    b.AbstractProductB.Init()
    b.IAbstractProductB = b
}

func (b *ZzzProductB) Action() {
    fmt.Printf("ZzzProductB Action\n")
}

func NewZzzProductB() *ZzzProductB {
    b := &ZzzProductB{}
    b.Init()
    return b
}
