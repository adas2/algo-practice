package practice

// LC # 69 easy

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	low, high := 0, x
	for low <= high {
		candidate := (low + high) / 2
		if candidate*candidate <= x &&
			(candidate+1)*(candidate+1) > x {
			return candidate
		} else if candidate*candidate > x {
			high = candidate - 1
		} else { // <= x
			low = candidate + 1
		}
	}
	return -1
}

// logic:
// sqrt(x) < x/2 start with x/2 use binary search to find candidate
// low=0, high=x, cand = (high+low)/2 i.e. x/2 so on.
