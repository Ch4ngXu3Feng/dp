package dp

import "fmt"

type IObject interface {
    Request()
}

type Object struct {
    IObject
}

func (o *Object) Init() {
    fmt.Printf("Object Init\n")
}

func (o *Object) Request() {
    fmt.Printf("Object Request\n")
}

func NewObject() *Object {
    o := &Object{}
    o.Init()
    return o
}

type ObjectProxy struct {
    IObject
    object *Object
}

func (p *ObjectProxy) Init() {
    fmt.Printf("ObjectProxy Init\n")
    p.IObject = p
}

func (p *ObjectProxy) load() {
    fmt.Printf("ObjectProxy load\n")
    if p.object == nil {
        p.object = NewObject()
    }
}

func (p *ObjectProxy) check() {
    fmt.Printf("ObjectProxy check\n")
}

func (p *ObjectProxy) Request() {
    p.load()
    p.check()
    p.object.Request()
}

func NewObjectProxy() *ObjectProxy {
    p := &ObjectProxy{}
    p.Init()
    return p
}
