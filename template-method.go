package dp

import "fmt"

type ITemplate interface {
    OperationXxx()
    OperationZzz()
}

type BaseTemplate struct {
    ITemplate
}

func (t *BaseTemplate) Operation() {
    t.OperationXxx()
    t.OperationZzz()
}

func (t *BaseTemplate) Init() {
    fmt.Printf("BaseTemplate Init\n")
    t.ITemplate = t
}

func NewBaseTemplate() *BaseTemplate {
    t := &BaseTemplate{}
    t.Init()
    return t
}

type ConcreteTemplate struct {
    BaseTemplate
}

func (t *ConcreteTemplate) OperationXxx() {
    fmt.Printf("ConcreteTemplate OperationXxx\n")
}

func (t *ConcreteTemplate) OperationZzz() {
    fmt.Printf("ConcreteTemplate OperationZzz\n")
}

func (t *ConcreteTemplate) Init() {
    fmt.Printf("ConcreteTemplete Init\n")
    t.BaseTemplate.Init()
    t.ITemplate = t
}

func NewConcreteTemplate() *ConcreteTemplate {
    t := &ConcreteTemplate{}
    t.Init()
    return t
}
