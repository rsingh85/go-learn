package main

import "fmt"

func main() {
  fmt.Println(getMessage("Foo"))
  fmt.Println(add(23, 32))
}

func getMessage(name string) string {
  return "Hello " + name
}

func add(x, y int) int {
  return x + y
}

