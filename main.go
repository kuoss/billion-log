package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 100000) // Semaphore to limit concurrency

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			for k := 0; k <= 999; k++ {
				wg.Add(1)
				sem <- struct{}{} // Acquire a slot
				go func(i, j, k int) {
					defer wg.Done()
					defer func() { <-sem }() // Release the slot
					fmt.Printf("I%03d-J%03d-K%03d %s\n", i, j, k, uuid.New())
				}(i, j, k)
			}
		}
	}

	wg.Wait() // Wait for all goroutines to finish
}
