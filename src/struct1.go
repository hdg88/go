package main

import (
  "fmt"
)

type Saiyan struct {
  Name string
  Power int
}

func main() {
  // multiple ways of initializing
  goku := Saiyan {
    Name: "Goku",
    Power: 9000,
  }
  fmt.Printf ("his name is %s and his power is %d", goku.Name, goku.Power)
}

//https://www.google.ca/?gws_rd=cr&dcr=0&ei=BLClWreaFI-d5wLEg7ugBQ
