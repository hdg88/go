package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}
//the type *Saiyan is the receiver of the Super method
func (s *Saiyan) Super() {
  s.Power+=10000
}

func main() {
  goku := &Saiyan {
    Name: "Goku",
    Power: 9000,
  }
  fmt.Printf("Power before calling Super: \t%d\n", goku.Power)
  goku.Super()
  fmt.Printf("Power after calling Super: \t%d\n", goku.Power)
}
