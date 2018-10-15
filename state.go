package dp

import "fmt"

type StateContext struct {
    State *State
}

func (c *StateContext) Init() {
    fmt.Printf("State Context Init\n")
}

func (c *StateContext) Request() {
    c.State.Handle()
}

func NewStateContext(s *State) *StateContext {
    c := &StateContext{
        State: s,
    }
    c.Init()
    return c
}

type IState interface {
    Handle()
}

type State struct {
    IState
}

func (s *State) Init() {
    s.IState = s
    fmt.Printf("State Init\n")
}

func NewState() *State {
    s := &State{}
    s.Init()
    return s
}

type XxxState struct {
    State
}

func (s *XxxState) Init() {
    s.State.Init()
    s.IState = s
    fmt.Printf("XxxState Init\n")
}

func (s *XxxState) Handle() {
    fmt.Printf("XxxState Handle\n")
}

func NewXxxState() *XxxState {
    s := &XxxState{}
    s.Init()
    return s
}

type ZzzState struct {
    State
}

func (s *ZzzState) Init() {
    s.State.Init()
    s.IState = s
    fmt.Printf("ZzzState Init\n")
}

func (s *ZzzState) Handle() {
    fmt.Printf("ZzzState Handle\n")
}

func NewZzzState() *ZzzState {
    s := &ZzzState{}
    s.Init()
    return s
}
