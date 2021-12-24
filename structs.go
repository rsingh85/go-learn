package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1, 2})

  v1 := Vertex{0, 0}
  v2 := Vertex{0, 0}

  fmt.Println(v1, v2, v1 == v2)
}
