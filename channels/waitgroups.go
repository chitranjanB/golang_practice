package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Printf("Started goroutine %d\n", i)
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
