package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroups(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main is done")
}
