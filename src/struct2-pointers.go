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
  // the &Saiyan gives the memory address of the initialized Saiyan variable.
  goku := &Saiyan {
    Name: "Goku",
    Power: 9000,
  }
  fmt.Printf("Address of goku: \t%p \n",goku)

  Super(goku)
  fmt.Println(goku.Power)
}
//*X means pointer to value of type X.
func Super (s *Saiyan) {
  fmt.Printf("Address of s: \t\t%p \n",s)
  s.Power+=10000
}
