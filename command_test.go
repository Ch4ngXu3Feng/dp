package dp

import (
    "testing"
    "fmt"
)

func TestCommand(*testing.T) {
    fmt.Printf("\nTest Command:\n\n")

    receiver := NewReceiver()
    command := NewXxxCommand(receiver)
    invoker := NewInvoker()
    invoker.AddCommand(&command.Command)
    invoker.Trigger()
}
