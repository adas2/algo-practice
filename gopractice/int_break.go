package practice

import "math"

// lic # 343 Integer Break

// Non DP O(1) solution: https://twchen.gitbook.io/leetcode/dynamic-programming/integer-break
func integerBreak2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if (n % 3) == 0 {
		return int(math.Pow(3, float64(n/3)))
	}
	// two 2s
	if (n % 3) == 1 {
		return 4 * int(math.Pow(3, float64((n-4)/3)))
	}
	// one 2
	return 2 * int(math.Pow(3, float64((n-2)/3)))
}

// simple: DP solution with T O(N^2), O(N) space
func integerBreak(n int) int {

	// remember the int breaks for each num till n
	nMap := make([]int, n+1)
	// base case
	nMap[2] = 1

	// number ranges
	for i := 3; i <= n; i++ {
		// int k ranges, note only half needs to covered, other half is repeated
		for x := 1; x <= i/2; x++ {
			m := maxInt(nMap[x]*(i-x), nMap[i-x]*x, x*(i-x))

			if m > nMap[i] {
				nMap[i] = m

				// fmt.Printf("x: %d, i-x: %d, max: %d map: %v\n", x, i-x, m, nMap[2:])
			}
		}
	}

	return nMap[n]
}

func maxInt(a, b, c int) int {
	var max int = a
	if b > a {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// logic:
// Intuition: for any number n, the k nums should be chosen such that they are similar/equal
// Product = (n-1)*1 << (n-x)*x wheere x >> 2
// generate a hash table with the optimal k for each number from 2->n
// for each num i, change x from 2 to n-2 and see:
// if P(i) ->(i-x)*x > current P(i) then choose x
// Worst case time: O(n*n/2); i.e. for each number try all k's from [2,n/2]
// e.g. 12=6+6, 6*6=34
// 12=4+4+4, 4^3=64
// 12=3+3+3+3. 3^4=81
// 12=2+2+2+2+2+2, 2^6=64
// 12-3=9, 9-3=6,
