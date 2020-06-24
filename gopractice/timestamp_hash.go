package practice

// Given a history hash with time stamp, find value of a key with the latest timestamp <= T
// [key, <val, timestamp>]

import (
	"fmt"
	"time"
)

const (
	CORRECT  = iota
	TOO_LOW  = iota
	TOO_HIGH = iota
)

// Entry type
type Entry struct {
	val int       //value
	ts  time.Time //timestamp
}

// FindPredecessorEvent return the latest value from history with timestamp <= t
func FindPredecessorEvent(hMap map[string][]Entry, key string, aTime time.Time) int {
	// Appy binary search
	elist, err := Get(hMap, key)
	if err != nil {
		return -1
	}
	// fmt.Println(elen)
	index := BinarySearch(elist, aTime, 0, len(elist)-1)
	fmt.Println("index", index)
	return elist[index].val
}

// Get entry
func Get(hMap map[string][]Entry, key string) ([]Entry, error) {
	if _, exists := hMap[key]; exists {
		return hMap[key], nil
	}
	return nil, fmt.Errorf("No entry for key %v", key)
}

func Put(hMap map[string][]Entry, key string, val int, t time.Time) error {
	// create entry
	e := Entry{val: val, ts: t}
	hMap[key] = append(hMap[key], e)
	return nil
}

func AttemptGuess(ts time.Time, guess time.Time) int {
	if ts.Equal(guess) {
		return CORRECT
	} else if ts.After(guess) {
		fmt.Println("low", guess.String())
		return TOO_LOW
	} else if ts.Before(guess) {
		fmt.Println("high", guess.String())
		return TOO_HIGH
	} else {
		return -1 //err
	}
}

// BinarySearch find the elemtnor closest element <= target
func BinarySearch(arr []Entry, target time.Time, low int, high int) int {
	if low > high {
		return -1 //error case
	}
	candidate := -1

	// begin iteration in while style in C
	for low <= high {
		mid := low + (high-low)/2
		fmt.Println("index", mid)
		guess := arr[mid].ts
		// if guess <= target
		if AttemptGuess(target, guess) == CORRECT || AttemptGuess(target, guess) == TOO_LOW {
			candidate = mid
			low = mid + 1
		} else { // AttemptGuess == TOO_HIGH
			high = mid - 1
		}
	}
	// no match found
	return candidate
}
