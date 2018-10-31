package dp

import (
    "fmt"
    "reflect"
)

type IAdapter interface {
    GetAdapter(interface{}) interface{}
}

type Adapter struct {
    IAdapter
}

func (a *Adapter) Init() {
    a.IAdapter = a
    fmt.Printf("Adapter Init\n")
}

func (a *Adapter) GetClassName(class interface{}) string {
    _type := reflect.TypeOf(class)
    name := _type.String()
    if len(name) == 0 {
        name = _type.Elem().String()
    }
    return name
}

func (a *Adapter) GetAdapter(class interface{}) interface{} {
    var r *Resource
    var f *File

    name := a.GetClassName(class)
    switch name {
    case a.GetClassName(r):
        return NewResource("TestResourceAdapter")

    case a.GetClassName(f):
        return NewFile("TestFileAdapter")
    }
    return nil
}

func NewAdapter() *Adapter {
    a := &Adapter{}
    a.Init()
    return a
}

type IResource interface {
    Show()
}

type Resource struct {
    Name string
    IResource
}

func (r *Resource) Show() {
    fmt.Printf("Resource Show: %s\n", r.Name)
}

func NewResource(name string) *Resource {
    r := &Resource{
        Name: name,
    }
    return r
}

type File struct {
    Name string
}

func (f *File) Show() {
    fmt.Printf("File Show: %s\n", f.Name)
}

func NewFile(name string) *File {
    f := &File{
        Name: name,
    }
    return f
}
