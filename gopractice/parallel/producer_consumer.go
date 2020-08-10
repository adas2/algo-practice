package main

import (
	"fmt"
	"sync"
	"time"
)

const buffSize = 10

// single producer
func producer(id int, itemBuff chan int) {
	for i := 0; i < buffSize; i++ {
		itemBuff <- i
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("prodcuer %d: created item %d\n", id, i)
	}
	close(itemBuff)
}

// single consumer
func consumer(id int, itemBuff chan int, done chan bool) {
	for v := range itemBuff {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("consumer %d: taking item %d\n", id, v)
	}
	done <- true
}

// main for generating producer consumer: single producer single consumer
func pcMain() {
	// define global buffer/channel to create/consume items
	itemBuff := make(chan int, buffSize)
	done := make(chan bool)

	go producer(0, itemBuff)
	go consumer(0, itemBuff, done)
	// wait till done
	<-done

}

// consumer using done channel
func consumer2(id int, itemBuff chan int, done chan bool) {
	for {
		select {
		case v, ok := <-itemBuff:
			if ok {
				fmt.Printf("consumer %d: taking item %d\n", id, v)
				time.Sleep(1 * time.Nanosecond)
			} else {
				// channel closed
				done <- true
				return
			}
		}
	}
}

// single producer multiple consumer
func mainSPMC() {
	itemBuff := make(chan int, buffSize)
	done := make(chan bool)

	go producer(0, itemBuff)

	for j := 0; j < 5; j++ {
		go consumer(j, itemBuff, done)
	}

	<-done

}

// multi producer
func producer3(id int, itemBuff chan int, w *sync.WaitGroup) {
	defer w.Done()
	for i := id * 10; i < id*10+10; i++ {
		itemBuff <- i
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("prodcuer %d: created item %d\n", id, i)
	}
}

// multi consumer using waitgroup
func consumer3(id int, itemBuff chan int, w *sync.WaitGroup) {
	for {
		select {
		case v, ok := <-itemBuff:
			if ok {
				fmt.Printf("consumer %d: taking item %d\n", id, v)
			} else {
				// channel closed
				w.Done()
				return
			}
		}
	}
}

// multi producer multi consumer
func mainMPMC() {
	// define global buffer/channel to create/consume items
	itemBuff := make(chan int, buffSize)

	wg1 := new(sync.WaitGroup)
	// start producers
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go producer3(i, itemBuff, wg1)
	}

	wg2 := new(sync.WaitGroup)
	// start consumers
	for j := 0; j < 3; j++ {
		wg2.Add(1)
		go consumer3(j, itemBuff, wg2)
	}

	// wait till all producers threads finish, and close buffer
	wg1.Wait()
	close(itemBuff)

	// wait for consumers to finish
	wg2.Wait()

}
