package main

import (
        "fmt"
        "runtime"
        "time"
)

func main() {
        // We can write multiple if statements.
        age := 42
        if age == 16 {
                fmt.Println("Too young")
        }
        if age == 42 {
                fmt.Println("Adult")
        }
        if age == 70 {
                fmt.Println("Senior")
        }

        // But switch is more readable.
        // [switch] [case]
        switch age {
        case 16:
                fmt.Println("Too young")
        case 42:
                fmt.Println("Adult")
        case 70:
                fmt.Println("Senior")
        }

        // Switch with a short statement.
        switch os := runtime.GOOS; os {
        case "linux":
                fmt.Println("Linux.")
        }

        t := time.Now()
        // Switch without a condition is the same as switch true.
        switch {
        case t.Hour() < 12:
                fmt.Println("Good morning!")
        }
}
