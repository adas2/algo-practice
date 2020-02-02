package main

import (
	"runtime"
	"time"
)

func testConcurrency() {

	runtime.GOMAXPROCS(2)

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
