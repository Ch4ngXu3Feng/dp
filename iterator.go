package dp

import "fmt"

type IIterator interface {
    First() interface{}
    Next() interface{}
    Last() interface{}
}

type Iterator struct {
    IIterator
}

func (i *Iterator) Init() {
    fmt.Printf("Iterator Init\n")
    i.IIterator = i
}

func (i *Iterator) First() interface{} {
    if _, ok := i.IIterator.(*Iterator); ok {
        panic("Iterator Next oops\n")
    }
    return i.IIterator.First()

}

func (i *Iterator) Next() interface{} {
    if _, ok := i.IIterator.(*Iterator); ok {
        panic("Iterator Next oops\n")
    }
    return i.IIterator.Next()
}

func (i *Iterator) Last() interface{} {
    if _, ok := i.IIterator.(*Iterator); ok {
        panic("Iterator Next oops\n")
    }
    return i.IIterator.Last()
}

func NewIterator() *Iterator {
    i := &Iterator{}
    i.Init()
    return i
}

type ElementsIterator struct {
    Iterator
    index int
    elements *Elements
}

func (i *ElementsIterator) Init() {
    fmt.Printf("ElementsIterator Init\n")
    i.Iterator.Init()
    i.IIterator = i
}

func (i *ElementsIterator) First() interface{} {
    i.index = 0
    return i.elements.Data[i.index]
}

func (i *ElementsIterator) Next() interface{} {
    if len(i.elements.Data) == i.index + 1 {
        return nil
    }
    i.index += 1
    return i.elements.Data[i.index]
}

func (i *ElementsIterator) Last() interface{} {
    return nil
}

func NewElementsIterator(elements *Elements) *ElementsIterator {
    i := &ElementsIterator{
        elements: elements,
    }
    i.Init()
    return i
}

type Elements struct {
    Data []string
}

func NewElements(data ...string) *Elements {
    elements := &Elements{
        Data: data,
    }
    return elements
}
