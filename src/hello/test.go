package main

import (
    "fmt"
    "time"
    "sync"
    "math/rand"
)


func boring(i int, wg* sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    fmt.Println("boring", i);
}

func main() {

    var wg sync.WaitGroup
    n := 100
    wg.Add(n)
    for i:=0; i<n; i++ {
        go boring(i, &wg)
    }
    wg.Wait()
}
