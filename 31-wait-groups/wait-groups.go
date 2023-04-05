package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
  fmt.Printf("Worker %d starting\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d done\n", id)
}

func main() {
  var waitGroup sync.WaitGroup

  for i := 1; i <= 5; i++ {
    waitGroup.Add(1) // increment the wait group counter

    i := i

    go func() {
      defer waitGroup.Done() // It will be executed after goes out of scope in this closure and also it will decrement the counter
      worker(i)
    }()
  }

  waitGroup.Wait() // Wait for all the goroutines to end
}
