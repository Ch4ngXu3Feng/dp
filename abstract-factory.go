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
    f.IAbstractFactory = f
    fmt.Printf("AbstractFactory Init\n")
}

func (f *AbstractFactory) CreateProductA() *AbstractProductA {
    panic("AbstractFactory CreateProductA called\n")
}

func (f *AbstractFactory) CreateProductB() *AbstractProductB {
    panic("AbstractFactory CreateProductB called\n")
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
    f.AbstractFactory.Init()
    f.IAbstractFactory = f
    fmt.Printf("XxxFactory Init\n")
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
    f.AbstractFactory.Init()
    f.IAbstractFactory = f
    fmt.Printf("ZzzFactory Init\n")
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
    a.IAbstractProductA = a
    fmt.Printf("AbstractProductA Init\n")
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
    a.AbstractProductA.Init()
    a.IAbstractProductA = a
    fmt.Printf("XxxProductA Init\n")
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
    a.AbstractProductA.Init()
    a.IAbstractProductA = a
    fmt.Printf("ZzzProductA Init\n")
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
    b.IAbstractProductB = b
    fmt.Printf("AbstractProductB Init\n")
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
    b.AbstractProductB.Init()
    b.IAbstractProductB = b
    fmt.Printf("XxxProductB Init\n")
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
    b.AbstractProductB.Init()
    b.IAbstractProductB = b
    fmt.Printf("ZzzProductB Init\n")
}

func (b *ZzzProductB) Action() {
    fmt.Printf("ZzzProductB Action\n")
}

func NewZzzProductB() *ZzzProductB {
    b := &ZzzProductB{}
    b.Init()
    return b
}
