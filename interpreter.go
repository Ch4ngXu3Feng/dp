package dp

import (
    "fmt"
    "strings"
)

type Context struct {
    Input string
}

type IExpression interface {
    Interpret(*Context) bool
}

type Expression struct {
    IExpression
}

func (e *Expression) Init() {
    fmt.Printf("Expression Init\n")
    e.IExpression = e
}

func NewExpression() *Expression {
    e := &Expression{}
    e.Init()
    return e
}

type TerminalExpression struct {
    Expression
    Data string
}

func (e *TerminalExpression) Init() {
    fmt.Printf("TerminalExpression Init\n")
    e.Expression.Init()
    e.IExpression = e
}

func (e *TerminalExpression) Interpret(c *Context) bool {
    return strings.Contains(c.Input, e.Data)
}

func NewTerminalExpression(data string) *TerminalExpression {
    e := &TerminalExpression{
        Data: data,
    }
    e.Init()
    return e
}

type AndExpression struct {
    Expression
    exp1 *Expression
    exp2 *Expression
}

func (e *AndExpression) Interpret(c *Context) bool {
    return e.exp1.Interpret(c) && e.exp2.Interpret(c)
}

func NewAndExpression(exp1 *Expression, exp2 *Expression) *AndExpression {
    e := &AndExpression{
        exp1: exp1,
        exp2: exp2,
    }
    e.Init()
    return e
}

type OrExpression struct {
    Expression
    exp1 *Expression
    exp2 *Expression
}

func (e *OrExpression) Interpret(c *Context) bool {
    return e.exp1.Interpret(c) || e.exp2.Interpret(c)
}

func NewOrExpression(exp1 *Expression, exp2 *Expression) *OrExpression {
    e := &OrExpression{
        exp1: exp1,
        exp2: exp2,
    }
    e.Init()
    return e
}
