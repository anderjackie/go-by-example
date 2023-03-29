package main

import "fmt"

// pings is a channel that only sends messages
func ping(pings chan<- string, msg string)  {
  pings <- msg // moves a string inside the channel - now pings has a message that it can sends
}

// pings is a channel that only receives messages
// pongs is a channel that only sends messages
func pong(pings <-chan string, pongs chan<- string)  {
  msg := <-pings // msg receives whats is stored inside pings
  pongs <- msg // pongs now has a message that can be send
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)

  ping(pings, "passed message")
  pong(pings, pongs)

  fmt.Println(<-pongs)
}

