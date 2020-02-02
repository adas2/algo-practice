package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ch := make(chan string, 10)
	quit := make(chan bool, 1)

	// pupulate channel
	input := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(input, " ")
	for _, w := range words {
		ch <- w
	}
	close(ch)

	// kill after sometime
	go func() {
		time.Sleep(5 * time.Second)
		quit <- true
	}()

loop:
	// drain channel
	for {
		select {
		case w, ok := <-ch:
			if ok {
				fmt.Println(w)
			}
		case <-quit:
			fmt.Println("Quitting..")
			break loop
		default:
			fmt.Println("Waiting")
		}
	}

	return
}
