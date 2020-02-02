package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

var input = []string{
	"new",
	"san",
	"chi",
	"lon",
}

func testConcurrentAPI() {

	startTime := time.Now()
	var numProcessed = 0
	runtime.GOMAXPROCS(2)

	for _, symbol := range input {
		// add concurrency by go thread
		go func(symbol string) {
			resp, _ := http.Get("http://www.metaweather.com/api/location/search/?query=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			var loc LocationResponse
			err := json.Unmarshal(body, &loc)
			if err != nil {
				fmt.Println(err)
			}
			for _, val := range loc {
				fmt.Printf("%s is a %s located at %s\n", val.Title, val.LocationType, val.LattLong)
			}
			numProcessed++
		}(symbol)

	}

	// wait till all input symbols processed
	for numProcessed < len(input) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
}

//
type LocationResponse []struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int64  `json:"woeid"`
	LattLong     string `json:"latt_long"`
}
