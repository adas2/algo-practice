package practice

// Rand7 is designed using given func rand5() [CCI]
// non-deterministic solution
func Rand7(rand5 func() int) int {

	// create a way of uniquely generating each number upto any range > 7
	// one way is 5 * rand5() + rand5()

	// above will generate nums 0-24 with equal probability
	// num % 7 will thus generate 0-6, but 0,1,2,3 with more porbability for nums 21-24
	// we can thus discard 21-24 cases and in an endless while loop till 0-20 is generated

	for {

		num := 5*rand5() + rand5()

		if num < 21 {
			return num % 7
		}
	}

	// note this can take forever (or very slow) in extreme circumstances

}
