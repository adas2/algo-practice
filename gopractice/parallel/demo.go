package main

import (
	"runtime"
	"time"
)

func testConcurrency() {

	// defines the number of logical cores
	// w/o this single core will execute all threads
	runtime.GOMAXPROCS(4)

	for i := 0; i < 10; i++ {
		go func() {
			println("Hello")
			time.Sleep(10 * time.Millisecond)
		}()

	}

	for i := 0; i < 10; i++ {
		go func() {
			println("Go")
			time.Sleep(10 * time.Millisecond)
		}()

	}

	time.Sleep(10 * time.Millisecond)
}
