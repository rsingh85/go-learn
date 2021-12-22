package main

import "fmt"

var b bool
var idx1, idx2 = 10, 11 // initialise inline

func main() {
  var i int
  b = true
  s := "hello"
  const Pi = 3.142
  fmt.Println(i, b, s, idx1, idx2, Pi)
}
