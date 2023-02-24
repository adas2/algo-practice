package practice

import (
	"fmt"
	"sort"
)

// LC # 332

func findItinerary(tickets [][]string) []string {
	itinerarySet := make(map[string][]bool)
	destMap := make(map[string][]string)

	for _, it := range tickets {
		// add to src->dest map similar to adjacency matrix
		if _, exists := destMap[it[0]]; exists {
			destMap[it[0]] = append(destMap[it[0]], it[1])
		} else {
			destMap[it[0]] = append([]string{}, it[1])
		}

		// add every ticket to itin set set
		// note tickets can have multiple same routes i.e. <JFK->SFO>,<JFK->SFO>
		itinerarySet[it[0]] = append(itinerarySet[it[0]], false)
	}
	// sort the dest in each src in order
	for _, v := range destMap {
		sort.SliceStable(v, func(i, j int) bool { return v[i] < v[j] })
	}

	route := []string{"JFK"}    //candidate for backtracking
	result := make([]string, 0) //final result when all tickets are covered
	start := "JFK"

	findItineraryUtil(itinerarySet, destMap, start, route, len(tickets), &result)

	return result
}

// recursive func for backtracking
func findItineraryUtil(iSet map[string][]bool, destMap map[string][]string, src string, route []string, numTickets int, itin *[]string) bool {
	// when all tickets are used i.e. route len is tickets+1
	if len(route) == numTickets+1 {
		fmt.Printf("The itineray found: %v\n", route)
		*itin = route
		return true
	}

	// fmt.Printf("DestMap: %v\n", destMap)
	// fmt.Printf("ISet: %v\n", iSet)
	if _, exists := destMap[src]; !exists {
		return false
	}

	for index := range destMap[src] {
		dst := destMap[src][index]
		// mark as visited if not
		if iSet[src][index] == false {
			iSet[src][index] = true
			// recursive
			ret := findItineraryUtil(iSet, destMap, dst, append(route, dst), numTickets, itin)
			// backtracking
			iSet[src][index] = false
			// tricky: this is needed to terminate recursive backtracking
			if ret {
				return true
			}
		}

	}

	return false
}

// Logic:
// Create set with all itineraries map[<src, dest>]:true, a map one with source as key and list(dest) as val
// E.g. map['JFK']: ['ATL', 'SFO']
// Start with key which has 'JFK', if multiple store in lexigographic order
// Dequeue from set once <src, dest> pair is used
// return when all itins are used up
