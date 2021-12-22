package main

import "fmt"

func main() {
  fmt.Println(getMessage("Ravi"))
}

func getMessage(name string) string {
  return "Hello " + name
}
