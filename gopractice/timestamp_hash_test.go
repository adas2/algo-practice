package practice

import (
	"time"
	"fmt"
	"testing"
)

func TestFindPredecessorEvent(t *testing.T){
	hist := make(map[string][]Entry)
	Put(hist, "Key1", 100, getCurrentTime())
	time.Sleep(100 * time.Millisecond)
	Put(hist, "Key1", 50, getCurrentTime())
	// Put(hist, "Key2", 200, getCurrentTime())

	test := getCurrentTime()
	fmt.Println(test.String())
	time.Sleep(100 * time.Millisecond)
	
	Put(hist, "Key1", 30, getCurrentTime())
	// Put(hist, "Key2", 10, getCurrentTime())

	fmt.Println("hist[Key1]")
	for _, p := range hist["Key1"] {
        fmt.Println(p.val, p.ts.String())
    }

	val := FindPredecessorEvent(hist, "Key1", test)
	fmt.Println("Value", val)
}

func getCurrentTime() time.Time{
	return time.Now()
}
