package main

import (
	"fmt"
	"time"
)

func main() {
  // buffers 5 requests in the channel
  requests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests <- i
  }
  close(requests)

  limiter := time.Tick(200 * time.Millisecond)

  for req := range requests {
    <-limiter // triggers every 200 miliseconds
    fmt.Println("request", req, time.Now())
  }

    fmt.Printf("\n")

  // buffers 3 time's in the channel 
  burstyLimiter := make(chan time.Time, 3)
  for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
  }

  // Routine to fill up the limiter channel every 200 miliseconds
  // But since the channel is full nothing is added to it
  go func ()  {
    for t := range time.Tick(200 * time.Millisecond) {
      burstyLimiter <- t
      fmt.Println(burstyLimiter)
    }
  }()

  burstyRequests := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    burstyRequests <- i
  }
  close(burstyRequests)

  // Go through the first 3 really quickly because the values are already there in the limiter channel
  for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
  }
}
