package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}

func NewSaiyan (name string, power int) Saiyan {
  return Saiyan {
    Name: name,
    Power: power,
  }
}

func main() {
  //this is how we instantiate using a constructor
  goku := NewSaiyan("Goku", 9000)
  //or, instead, we can instantiate it this easier way.
  gohan := &Saiyan {
    Name: "Gohan",
    Power: 1000,
  }

  fmt.Printf("Power of Goku: \t%d\n", goku.Power)
  fmt.Printf("Name: \t%s\n", goku.Name)
  fmt.Printf("Power of Gohan: \t%d\n", gohan.Power)
  fmt.Printf("Name: \t%s\n", gohan.Name)}
