package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
  fmt.Println("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}

func main() {
  done := make(chan bool, 1)

  go worker(done)

  // This blocks the execution of the program until done receives a value
  // If this block operation is not here the code will execute the worker goroutine and finish so th result will never be seen
  <-done
}
