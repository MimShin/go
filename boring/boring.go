package main

import (
  "fmt"
)

func main() {
  n1, n2 := 10, 10

  c1 := goBoring("Rosha", n1)
  c2 := goBoring("Shana", n2)

  for i:=0; i<n1+n2; i++ {
    select {
    case x := <-c1:
      fmt.Println(x)
    case x := <-c2:
      fmt.Println(x)
    }
  }
}

func goBoring(name string, count int) <-chan string {
  c := make(chan string)  
  go boring(name, count, c)
  return c
}

func boring(name string, count int, c chan string) {
  for i:=0;i<count;i++ {
    c <- fmt.Sprintf("%s: %d", name, i)
  }
}
