package strs

import (
	"strconv"
)

// EPI. pp 84
// Generate n-thsequence in the look and say str{}
// {1, 11, 21, 1211, 111221, 312211, 13112221, ....}

// var lookup = []string{"zero", "one", "two", "three", }

func lookAndSaySequence(n int) string {

	// iterative map
	iMap := make(map[int]string)

	for i := 1; i <= n; i++ {
		if i == 1 {
			iMap[i] = "1"
			continue
		}
		// else populate next on
		iMap[i] = genLookAndSay(iMap[i-1])
	}

	return iMap[n]
}

// util to generate the string
func genLookAndSay(s string) string {

	var res string

	cnt := 1
	i := 1
	for i < len(s) {
		if s[i] == s[i-1] {
			cnt++
			// continue
		} else {
			// s[i] != s[i-1]
			res = res + strconv.Itoa(cnt) + string(s[i-1])
			// reset cnt
			cnt = 1
		}

		i++
	}
	// when i = len(s)
	res += strconv.Itoa(cnt) + string(s[i-1])

	// fmt.Printf("gen %s string: %s\n", s, res)
	return res
}

// to generate n-th str, use (n-1)th str
// count subsequent same nums till s[i] != s[i-1]
// gen n*"s[i-1]" in result
