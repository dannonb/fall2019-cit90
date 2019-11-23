package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var counter int 

func main() {
	q := 100
	counter := 0
	wg.Add(q)
	for i := 0; i < q; i++ {
		go foo(i)
		counter++
		defer mu.Lock()
		
	}
	wg.Wait()
	fmt.Printf("counter: %v", counter)
}

func foo(n int) {
	fmt.Println(n)
	wg.Done()
}

