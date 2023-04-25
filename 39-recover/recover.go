package main

import "fmt"

func mayPanic() {
  panic("a problem")
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered. Error:\n", r)
    }
  }()

  mayPanic()

  fmt.Println("after mayPanic()") // It never gets executed because the program panicked and resumed in the deferred function
}
