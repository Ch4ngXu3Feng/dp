package dp

import "fmt"

type StrategyContext struct {
    Strategy *Strategy
}

func (c *StrategyContext) Init() {
    fmt.Printf("Strategy Context Init\n")
}

func (c *StrategyContext) SetStrategy(name int) {
    switch name {
    case 0:
        c.Strategy = &NewXxxStrategy().Strategy
    case 1:
        c.Strategy = &NewZzzStrategy().Strategy
    }
}

func (c *StrategyContext) Do() {
    c.Strategy.Process()
}

func NewStrategyContext() *StrategyContext {
    c := &StrategyContext{}
    c.Init()
    return c
}

type IStrategy interface {
    Process()
}

type Strategy struct {
    IStrategy
}

func (s *Strategy) Init() {
    fmt.Printf("Strategy Init\n")
    s.IStrategy = s
}

func (s *Strategy) Process() {
    fmt.Printf("Strategy Process\n")
    s.IStrategy.Process()
}

func NewStrategy() *Strategy {
    s := &Strategy{}
    s.Init()
    return s
}

type XxxStrategy struct {
    Strategy
}

func (s *XxxStrategy) Init() {
    fmt.Printf("XxxStrategy Init\n")
    s.Strategy.Init()
    s.IStrategy = s
}

func (s *XxxStrategy) Process() {
    fmt.Printf("XxxStrategy Process\n")
}

func NewXxxStrategy() *XxxStrategy {
    s := &XxxStrategy{}
    s.Init()
    return s
}

type ZzzStrategy struct {
    Strategy
}

func (s *ZzzStrategy) Init() {
    s.Strategy.Init()
    s.IStrategy = s
    fmt.Printf("ZzzStrategy Init\n")
}

func (s *ZzzStrategy) Process() {
    fmt.Printf("ZzzStrategy Process\n")
}

func NewZzzStrategy() *ZzzStrategy {
    s := &ZzzStrategy{}
    s.Init()
    return s
}
