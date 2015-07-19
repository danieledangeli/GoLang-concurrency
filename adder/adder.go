package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

func increment(p *int64) {
	atomic.AddInt64(p, 1)
}

func main() {
	var shared int64 = 0

	//wait Group for both thread to finish
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000000; i++ {
			increment(&shared)
		}

		defer wg.Done()

	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			increment(&shared)
		}

		defer wg.Done()
	}()

	//wait
	wg.Wait()

	fmt.Println(shared);
}