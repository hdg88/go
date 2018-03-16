package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
  Father *Saiyan
}



func main() {
  //this is how we instantiate using a constructor
  gohan := &Saiyan{
    Name: "Gohan",
    Power: 1000,
    Father: &Saiyan {
      Name: "Goku",
      Power: 9001,
      Father: nil,
    },
  }


  fmt.Printf("Power of Goku: \t%d\n", gohan.Father.Power)
  fmt.Printf("Name: \t%s\n", gohan.Father.Name)
  fmt.Printf("Power of Gohan: \t%d\n", gohan.Power)
  fmt.Printf("Name: \t%s\n", gohan.Name)
}
