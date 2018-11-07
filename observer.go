package dp

import "fmt"

type ISubject interface {
    Attach(*Observer)
    Notify()
    GetState() int
    SetState(int)
}

type Subject struct {
    ISubject
    state int
    observers []*Observer
}

func (s *Subject) Init() {
    fmt.Printf("Subject Init\n")
    s.ISubject = s
}

func (s *Subject) Attach(o *Observer) {
    s.observers = append(s.observers, o)
}

func (s *Subject) Notify() {
    for _, o := range s.observers {
        o.Update(s)
    }
}

func (s *Subject) GetState() int {
    return s.state
}

func (s *Subject) SetState(state int) {
    s.state = state
}

func NewSubject() *Subject {
    s := &Subject{}
    s.Init()
    return s
}

type ConcreteSubject struct {
    Subject
}

func (s *ConcreteSubject) Init() {
    fmt.Printf("ConcreteSubject Init\n")
    s.Subject.Init()
    s.ISubject = s
}

func (s *ConcreteSubject) SetState(state int) {
    s.Subject.SetState(state)
    s.Notify()
}

func NewConcreteSubject() *ConcreteSubject {
    s := &ConcreteSubject{}
    s.Init()
    return s
}

type IObserver interface {
    Update(*Subject)
}

type Observer struct {
    IObserver
}

func (o *Observer) Init() {
    fmt.Printf("Observer Init\n")
    o.IObserver = o
}

func (o *Observer) Update(s *Subject) {
    if _, ok := o.IObserver.(*Observer); ok {
        panic("Observer Update oops\n")
    }
    o.IObserver.Update(s)
}

func NewObserver() *Observer {
    o := &Observer{}
    o.Init()
    return o
}

type XxxObserver struct {
    Observer
}

func (o *XxxObserver) Init() {
    fmt.Printf("XxxObserver Init\n")
    o.Observer.Init()
    o.IObserver = o
}

func (o *XxxObserver) Update(s *Subject) {
    fmt.Printf("XxxObserver: Subject state is %d\n", s.GetState())
}

func NewXxxObserver() *XxxObserver {
    o := &XxxObserver{}
    o.Init()
    return o
}

type ZzzObserver struct {
    Observer
}

func (o *ZzzObserver) Init() {
    fmt.Printf("ZzzObserver Init\n")
    o.Observer.Init()
    o.IObserver = o
}

func (o *ZzzObserver) Update(s *Subject) {
    fmt.Printf("ZzzObserver: Subject state is %d\n", s.GetState())
}

func NewZzzObserver() *ZzzObserver {
    o := &ZzzObserver{}
    o.Init()
    return o
}
