package practice

import (
	"math"
)

// some arbitrary threshold for floating point diff
const threshold float64 = 0.000001

// IsPowerTen :
// Given a num find if it is power of 10
func IsPowerTen(num float64) bool {
	// corner case
	if num == 0 {
		return false
	}

	// define temp vars
	var factor float64 = 1

	if num >= 1 {
		// multiple factor and check if equals number
		for num >= factor {
			if factor == num {
				return true
			}
			factor *= 10
		}

	} else {
		// divide factor check if equals number
		for num <= factor {
			if IsEqualFloat(num, factor) {
				return true
			}
			factor /= 10
		}

	}

	return false
}

// IsEqualFloat approximates equality for floating point numbers
func IsEqualFloat(a, b float64) bool {
	if diff := math.Abs(a - b); diff < threshold {
		return true
	}

	return false
}
