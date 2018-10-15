package dp

import (
    "fmt"
)

type IAdapter interface {
    String() string
}

type Adapter struct {
    i int
}

func (a *Adapter) Init() {
    fmt.Printf("State Context Init\n")
}

func (a *Adapter) String() string {
    return "11"
}

func NewAdapter() *Adapter {
    a := &Adapter{}
    a.Init()
    return a
}

type ISubject interface {
}

type Subject struct {
    ISubject
}

func (s *Subject) Init() {
    s.ISubject = s
}

func NewSubject() *Subject {
    s := &Subject{}
    s.Init()
    return s
}
