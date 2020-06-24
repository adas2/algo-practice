package practice

// ConvertToRoman :
// Given an integer convert to Roman numernals
func ConvertToRoman(num int) string {

	// init variables
	var outp, partial string

	// recursively call helper func
	ConvertToRomanHelper(num, partial, &outp)

	return outp
}

// ConvertToRomanHelper using recursive calls
func ConvertToRomanHelper(num int, partial string, outp *string) {

	// fmt.Println(num, partial)
	// terminating condition
	if num == 0 {
		*outp = partial
		return
	}
	// find the largest int below given num
	val, sym := FindLargestBelowNum(num)
	// fmt.Printf("largets below: %d sym: %s\n", val, sym)

	// how many time sym will appear
	factor := num / val
	residue := num % val

	// fmt.Println(factor, residue)

	for i := 0; i < factor; i++ {
		partial += sym
	}
	// recurse
	ConvertToRomanHelper(residue, partial, outp)

}

// FindLargestBelowNum finds the sym and the num just given the given num
// if eqaul return correct symbol
func FindLargestBelowNum(num int) (int, string) {
	// I             1
	// V             5
	// X             10
	// L             50
	// C             100
	// D             500
	// M             1000

	// IV 			 4
	// IX 			 9
	// XL  	         40
	// XC 			 90
	// CD 			 400
	// CM 			 900

	// int map
	ints := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	// string map
	syms := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	// sanity check
	if len(ints) != len(syms) {
		return -1, ""
	}

	var index int
	// linear traversal is good enough because of const size
	for i, v := range ints {

		// linearly sorted, hence break we hit larger value
		if v > num {
			break
		}

		index = i
	}

	return ints[index], syms[index]
}
