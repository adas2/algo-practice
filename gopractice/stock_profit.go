package practice

import "math"

// leetcode #122 and #121 (EPI)

type period struct {
	start int
	end   int
}

// multi transaction max profit
// seeling allowed only after buying
func maxProfit(prices []int) int {
	// max profit
	var maxP int
	// priceDiff array
	priceDiff := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		priceDiff[i] = prices[i] - prices[i-1]
	}

	// traverse for positive seuqences in priceDiff and add them up
	for _, val := range priceDiff {
		if val > 0 {
			maxP += val
		}

	}

	return maxP
}

// single transaction MAx profit
// selling only after buying
// Note: losses not captured, i.e. o/p 0 or more
func oneMaxProfit(prices []int) int {
	// candidate minimum and maxm, init to day1 prices
	var cMin, cMax int
	// buying and selling day index
	var buyDay, sellDay int
	var maxProfit int //set to INT_MIN

	// init with math
	cMin = math.MaxInt32
	cMax = math.MinInt32

	// O(n) traversal
	for i, val := range prices {
		// new min encountered
		if val < cMin {

			// Check if max profit exceeded with oldMin
			if (sellDay > buyDay) && maxProfit < (cMax-cMin) {
				maxProfit = cMax - cMin
			}

			// update min
			cMin = val
			buyDay = i
			// reset max
			cMax = val

		} else if val > cMax { //new nax encountered
			// update max
			cMax = val
			sellDay = i
		} // else val == cMax || val == cMin no extra gain/loss
	}
	// Check if max profit exceeded
	if (sellDay > buyDay) && maxProfit < (cMax-cMin) {
		maxProfit = cMax - cMin
	}

	return maxProfit
}

// logic:
//  e.g. input price: [7,1,5,3,6,4]
// find all the diff deltas i.e. d/dx store the priceDiff array,
// check for positive sequences and add them
