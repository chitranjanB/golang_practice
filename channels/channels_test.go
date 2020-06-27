package main

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
