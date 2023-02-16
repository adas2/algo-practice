package strs

// LC # 249:

// Use encoding of string using the relative distance between the letters
func groupStrings(strings []string) [][]string {
	// declare map
	grMap := make(map[string][]string)

	// traverse the string
	for _, s := range strings {
		// insert in map based on hash key
		key := genHashKey(s)
		grMap[key] = append(grMap[key], s)
	}

	// rearrange to 2d slice
	res := [][]string{}

	for _, strList := range grMap {
		res = append(res, strList)
	}

	return res
}

// util to generate hash key of a string using relative dist
func genHashKey(str string) string {
	// var dist rune = 'a'
	dist := make([]rune, len(str)-1)

	r := []rune(str)
	for i := 1; i < len(str); i++ {
		// note: dist slice is 1 less in len that original str
		// also 'a' is added to convert num to rune
		dist[i-1] = (r[i]-r[i-1]+26)%26 + 'a'
	}
	// fmt.Printf("hash key: %s\n", string(dist))
	return string(dist)
}

// Logic:
// E.g. "abc" relative dist --> (b-a) = 1, (c-b) = 1, i.e. <11>
// "acg" --> <23> convert it into str using + 'a'
// using above as hash key, populate map with strings with same key
// re-arrange the map into a 2-d string slice

// T: O(N*k) S = O(N)
