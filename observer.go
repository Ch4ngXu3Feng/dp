package dp

import "fmt"

type ISubject interface {
    Attach(*Observer)
    Notify()
    GetState() int
    SetState(int)
}

type Subject struct {
    state int
    observers []*Observer
    ISubject
}

func (s *Subject) Init() {
    s.ISubject = s
    fmt.Printf("Subject Init\n")
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
    s.Subject.Init()
    s.ISubject = s
    fmt.Printf("ConcreteSubject Init\n")
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
    o.IObserver = o
    fmt.Printf("Observer Init\n")
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
    o.Observer.Init()
    o.IObserver = o
    fmt.Printf("XxxObserver Init\n")
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
    o.Observer.Init()
    o.IObserver = o
    fmt.Printf("ZzzObserver Init\n")
}

func (o *ZzzObserver) Update(s *Subject) {
    fmt.Printf("ZzzObserver: Subject state is %d\n", s.GetState())
}

func NewZzzObserver() *ZzzObserver {
    o := &ZzzObserver{}
    o.Init()
    return o
}
