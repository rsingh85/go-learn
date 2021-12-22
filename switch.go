package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go is running on ")

  os := runtime.GOOS
 
  switch os {
    case "darwin":
      fmt.Println("OS X")
    case "linux":
      fmt.Println("Linux")
    default:
      fmt.Println("%s\n", os)
  } 
}
