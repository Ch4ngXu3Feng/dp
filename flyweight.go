package dp

import "fmt"

type IFlyweight interface {
    Operation()
}

type Flyweight struct {
    IFlyweight
}

func (f *Flyweight) Init() {
    fmt.Printf("Flyweight Init\n")
    f.IFlyweight = f
}

func (f *Flyweight) Operation() {
    if _, ok := f.IFlyweight.(*Flyweight); ok {
        panic("Flyweight Operation oops\n")
    }
    f.IFlyweight.Operation()
}

type XxxFlyweight struct {
    Flyweight
}

func (f *XxxFlyweight) Init() {
    fmt.Printf("XxxFlyweight Init\n")
    f.Flyweight.Init()
    f.IFlyweight = f
}

func (f *XxxFlyweight) Operation() {
    fmt.Printf("XxxFlyweight Operation\n")
}

func NewXxxFlyweight() *XxxFlyweight {
    f := &XxxFlyweight{}
    f.Init()
    return f
}

type ZzzFlyweight struct {
    Flyweight
}

func (f *ZzzFlyweight) Init() {
    fmt.Printf("ZzzFlyweight Init\n")
    f.Flyweight.Init()
    f.IFlyweight = f
}

func (f *ZzzFlyweight) Operation() {
    fmt.Printf("ZzzFlyweight Operation\n")
}

func NewZzzFlyweight() *ZzzFlyweight {
    f := &ZzzFlyweight{}
    f.Init()
    return f
}

type FlyweightManager struct {
    flyweights map[string]*Flyweight
}

func (m *FlyweightManager) Init() {
    fmt.Printf("FlyweightManager Init\n")
}

func (m *FlyweightManager) GetFlyweight(key string) *Flyweight {
    f, ok := m.flyweights[key]
    if !ok {
        switch key {
        case "Xxx":
            f = &NewXxxFlyweight().Flyweight
        case "Zzz":
            f = &NewZzzFlyweight().Flyweight
        }
        m.flyweights[key] = f
    }
    return f
}

func NewFlyweightManager() *FlyweightManager {
    m := &FlyweightManager{
        flyweights: map[string]*Flyweight{},
    }
    m.Init()
    return m
}
