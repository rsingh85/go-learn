package main

import "fmt"

func main() {
  fmt.Println(getMessage("Foo"))
  fmt.Println(add(23, 32))
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}

func getMessage(name string) string {
  return "Hello " + name
}

func add(x, y int) int {
  return x + y
}

// multiple results function
func swap(a, b string) (string, string) {
  return b, a
}
