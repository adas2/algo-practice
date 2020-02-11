package main

import (
	"fmt"
	"sync"
	"time"
)

func testMutex() {
	mtx := new(sync.Mutex)

	// some global variable changing
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mtx.Lock()
			go func() {
				fmt.Printf("i=%d, j=%d, i+j=%d\n", i, j, i+j)
				mtx.Unlock()
			}()
		}

		time.Sleep(10 * time.Millisecond)
	}
}

// go channel implementation of mutex
func testChanMutex() {
	ch := make(chan bool, 1)

	// some global variable changing
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ch <- true
			go func() {
				fmt.Printf("i=%d, j=%d, i+j=%d\n", i, j, i+j)
				<-ch
			}()
		}

		time.Sleep(10 * time.Millisecond)
	}
}
