package dp

import (
    "testing"
    "fmt"
)

func TestInterpreter(*testing.T) {
    fmt.Printf("Test Interpreter:\n\n")

    male := NewOrExpression(
        &NewTerminalExpression("Robert").Expression,
        &NewTerminalExpression("John").Expression,
    )

    married := NewAndExpression(
        &NewTerminalExpression("Julie").Expression,
        &NewTerminalExpression("Married").Expression,
    )

    ok1 := male.Interpret(&Context{
        Input: "John",
    })

    ok2 := married.Interpret(&Context{
        Input: "Married Julie",
    })

    fmt.Printf(
        "John is male? %t\nJulie is a married women? %t\n",
        ok1, ok2,
    )

    fmt.Printf("\n")
}
