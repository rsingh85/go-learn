package main

import "fmt"

func main() {
  var planets [8]string
  planets[0] = "earth"
  planets[1] = "jupiter"
  planets[2] = "saturn"
  planets[3] = "neptune"
  planets[4] = "mars"
  planets[5] = "uranus"
  planets[6] = "venus"
  planets[7] = "mercury"

  fmt.Println(planets)
  fmt.Println(len(planets))

  for i := 0; i < len(planets); i++ {
    fmt.Println(planets[i])
  }
}
