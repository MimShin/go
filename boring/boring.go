package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	n := 10
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	routines := make(map[int]bool)
	for i:=0; i<n; i++ {
		wg.Add(1)
		// WaitGroup and Mutex objects must be passed as address otherwise new objects will be created and they don't work
		go boring(i, routines, &wg, &mu)
	}
	wg.Wait()
}

func boring(n int, routines map[int]bool, wg* sync.WaitGroup, mu* sync.Mutex) {

	mu.Lock()	// Concurrent access problem without Mutex
	routines[n] = false
	mu.Unlock()

	fmt.Printf("Boring %d\n", n)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	wg.Done()

	mu.Lock()
	routines[n] = true;
	mu.Unlock()
}
