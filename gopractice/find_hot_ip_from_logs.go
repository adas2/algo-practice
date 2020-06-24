package practice

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
)

// FindHotIPAddress :
// Given an input stream of logs find the hottest IP addresses
func FindHotIPAddress(r io.Reader) string {
	// Read bytes buffer from inout
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

	// simple version
	x, err := regexp.Compile("[0-9]+.[0-9]+.[0-9]+.[0-9]+")

	// find all occurences
	ips := x.FindAll(buf, -1)
	fmt.Printf("All occurence of IP: %q\n", ips)

	// create Map with ip addresses
	cMap := CreateIPAddressMap(ips)

	// traverse map t find the hot address with max count
	var max int = -1
	var result string

	for k := range cMap {
		if max < cMap[k] {
			max, result = cMap[k], k
		}

	}

	return result
}

// CreateIPAddressMap creates a map of the IP addresses accessed with count
func CreateIPAddressMap(ips [][]byte) map[string]int {

	cMap := map[string]int{}

	for _, ip := range ips {
		if _, ok := cMap[string(ip)]; ok {
			//present
			cMap[string(ip)]++
		} else {
			//not present
			cMap[string(ip)] = 1
		}
	}

	return cMap
}
