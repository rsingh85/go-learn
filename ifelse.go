package main

import "fmt"

func main() {
  name := "Hello"
  if len(name) == 5 {
    fmt.Println("name is of length 5")
  } else {
    fmt.Println("name is not of length 5")
  }
}
