package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch := make(chan int)
	go generate(ch)

	// continue to filter
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		//flip channels to read from o/p channel
		ch = ch1
	}
}

// endless loop generating ints
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
	// close channel to prevent deadlock
	// close(ch)
}

// filters numbers from input channel based on given prime factor
func filter(in, out chan int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}
