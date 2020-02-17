package main

import (
	"fmt"
)

type Button struct {
	eventMap map[string][]chan string
}

func createButton() *Button {
	// return new(Button)
	b := new(Button)
	b.eventMap = make(map[string][]chan string)
	return b
}

// AddEventHandler adds a new handler corresponding to an event
func (b *Button) AddEventHandler(event string, handler chan string) {
	if _, present := b.eventMap[event]; present {
		b.eventMap[event] = append(b.eventMap[event], handler)
	} else {
		// not present
		b.eventMap[event] = [](chan string){handler}
		// b.eventMap[event] = append(b.eventMap[event], handler)
	}
}

// DeleteEventHandler deletes an exisitng handler
func (b *Button) DeleteEventHandler(event string, handler chan string) {
	if _, present := b.eventMap[event]; !present {
		fmt.Printf("Error: Event %s not present\n", event)
	} else {
		// present
		for idx, h := range b.eventMap[event] {
			if h == handler {
				// delete idx
				b.eventMap[event] = append(b.eventMap[event][:idx], b.eventMap[event][idx+1:]...)
				break
			}
		}
	}

}

// TriggerEvent triggers an event by calling the handlers
func (b *Button) TriggerEvent(event string, action string) {
	// fmt.Printf("Current Event %s:\n", event)
	if _, present := b.eventMap[event]; present {
		// push to all handlers for that event
		for _, h := range b.eventMap[event] {
			go func(handler chan string) {
				handler <- action
			}(h)
		}
	}

}

// create event using channels
func testEvent() {
	myButton := createButton()

	// define event handler channels
	// NOTE: unbuffered channel is ok here since
	// we don't need buffer to keep the button from blokcing
	// since it publishes it's events though goroutines
	h1 := make(chan string)
	h2 := make(chan string)
	h3 := make(chan string)

	myButton.AddEventHandler("click", h1)
	myButton.AddEventHandler("click", h2)
	myButton.AddEventHandler("double click", h3)

	myButton.TriggerEvent("click", "call clikcing function")

	// go threads start here
	go func() {
		for {
			if msg, ok := <-h1; ok {
				fmt.Printf("Event hanlder h1: %s\n", msg)
				// break
			}
		}
	}()

	go func() {
		for {
			if msg, ok := <-h2; ok {
				fmt.Printf("Event hanlder h2: %s\n", msg)
				// break
			}
		}
	}()

	go func() {
		for {
			if msg, ok := <-h3; ok {
				fmt.Printf("Event hanlder h3: %s\n", msg)
				// break
			}
		}
	}()

	myButton.DeleteEventHandler("click", h1)
	myButton.TriggerEvent("click", "call clicking function again")
	myButton.TriggerEvent("double click", "call open function")

	// sleep for some time to let go threads finish their work
	// time.Sleep(10 * time.Millisecond)

	// wait till user input
	fmt.Scanln()
}
