package main

import (
  "fmt"
  "golang.org/x/tour/tree"
)

func main() {
  t1 := tree.New(1)
  t2 := tree.New(1)
  fmt.Println(t1, " - ", t2)
  fmt.Println(same(t1, t2))

  /*
  c1 := make(chan int)
  go walk(c1, tree.New(1), true)
  for {
    x, ok := <- c1
    if !ok {
      break
    }
    fmt.Println(x)
  }
  */
}

func walk(c chan int, t *tree.Tree, close_c bool) {
  if t != nil {
    walk(c, t.Left, false)
    c <- t.Value
    walk(c, t.Right, false)
  }

  if close_c {
    close(c)
  }
}

func goWalk(t *tree.Tree) <- chan int {
  c := make(chan int)
  go walk(c, t, true)
  return c
}

func same(t1, t2 *tree.Tree) bool {

  c1 := goWalk(t1)
  c2 := goWalk(t2)

  for {
    x1, ok1 := <- c1
    x2, ok2 := <- c2

    if !ok1 && !ok2 {
      return true
    }

    if !ok1 || !ok2 {
      return false
    }

    if x1 != x2 {
      return false
    }
  }
  return true
}
