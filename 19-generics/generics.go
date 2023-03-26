package main

import "fmt"

// Receives a map of any type K and V
// K has the comparable constraint which means that it can be compared with "!=" or "==" operators
// V has the any constraint which means that can be any type ov value (no constraint)
func MapKeys[K comparable, V any](m map[K]V) []K {
  r := make([]K, 0, len(m)) // create a slice of type K with size of length of the map

  for k := range m { // take a range of the map and ppends the keys to the slice
    r = append(r, k)
  }

  return r
}

type List[T any] struct {
  head, tail *element[T]
}

type element[T any] struct {
  next *element[T]
  val  T
}

func (lst *List[T]) Push(v T) {
  if lst.tail == nil {
    lst.head = &element[T]{val: v}
    lst.tail = lst.head
  } else {
    lst.tail.next = &element[T]{val: v}
    lst.tail = lst.tail.next
  }
}

func (lst *List[T]) GetAll() []T {
  var elems []T

  for e := lst.head; e != nil; e = e.next {
    elems = append(elems, e.val)
  }

  return elems
}

func main() {
  var m = map[int]string{1: "2", 2: "4", 4: "8"}

  fmt.Println("keys:", MapKeys(m))

  _ = MapKeys[int, string](m)

  lst := List[int]{}
  lst.Push(10)
  lst.Push(13)
  lst.Push(23)
  fmt.Println("list:", lst.GetAll())
}
