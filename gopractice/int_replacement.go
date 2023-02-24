package practice

// LC # 397

func integerReplacement(n int) int {

	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		//i is odd
		return Min(integerReplacement(n/2), integerReplacement(n/2+1)) + 2
	}

}

// logic:
// if even always divide n = n/2 --> 1 move
// if odd lookup min(result[n-1], result[n+1]), but this will be inefficient for large n
// to reduce the number of stack space needed for recursion, use following
// if n is odd, (n-1)/2 and (n+1)/2 can either be n/2 and n/2+1
//  so result = min of these two cases + 2 ( 1 for +/- , 1 for /2)
//  Approx complexity T=O(logn), S=O(logn) stack space needed
