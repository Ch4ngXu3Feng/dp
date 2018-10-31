package dp

import "fmt"

type Invoker struct {
    Command *Command
}

func (i *Invoker) Init() {
    fmt.Printf("Invoker Init\n")
}

func (i *Invoker) AddCommand(c *Command) {
    i.Command = c
}

func (i *Invoker) Trigger() {
    if i.Command != nil {
        i.Command.Execute()
    }
}

func NewInvoker() *Invoker {
    i := &Invoker{}
    i.Init()
    return i
}

type ICommand interface {
    Execute()
}

type Command struct {
    ICommand
}

func (c *Command) Init() {
    fmt.Printf("Command Init\n")
    c.ICommand = c
}

func NewCommand() *Command {
    c := &Command{}
    c.Init()
    return c
}

type XxxCommand struct {
    Receiver *Receiver
    Command
}

func (c *XxxCommand) Init() {
    c.Command.Init()
    c.ICommand = c
    fmt.Printf("XxxCommand Init\n")
}

func (c *XxxCommand) Execute() {
    c.Receiver.Action()
    fmt.Printf("XxxCommand Execute\n")
}

func NewXxxCommand(r *Receiver) *XxxCommand {
    c := &XxxCommand{
        Receiver: r,
    }
    c.Init()
    return c
}

type Receiver struct {
}

func (r *Receiver) Action() {
    fmt.Printf("Receiver Action\n")
}

func (r *Receiver) Init() {
    fmt.Printf("Receiver Init\n")
}

func NewReceiver() *Receiver {
    r := &Receiver{}
    r.Init()
    return r
}
