package dp

// easy problem
// find distance beween two strings
func HammingDistance(s1, s2 string) int {
	// length must match
	if len(s1) != len(s2) {
		return -1
	}

	var dist int = 0
	// traverse the strings left to right
	for i, j := 0, 0; i < len(s1); i, j = i+1, j+1 {
		if s1[i] != s2[j] {
			dist++
		}
	}
	return dist
}

// harder problem
// edit distance: changes (insert, remove, alter) needed to convert string1 --> string2
func EditDistance(s1, s2 string) int {

	// declare a 2D slice initialized to zero by default
	// eDist[i][j] represents the edit distance for substr s1 len i, s2 len j
	// eDist := [len(s1) + 1][len(s2) + 1]uint8{}

	eDist := make([][]int, len(s1)+1)
	for k := range eDist {
		eDist[k] = make([]int, len(s2)+1)
	}
	// case where one string is zero len, edit dis for other is len of str
	for i := range s1 {
		eDist[i][0] = i
	}

	for j := range s2 {
		eDist[0][j] = j
	}

	// for non zero lens using DP
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {

			if s1[i-1] == s2[j-1] {
				eDist[i][j] = eDist[i-1][j-1]
			} else { // s1[i] != s2[j]
				eDist[i][j] = 1 + findMin(
					eDist[i-1][j],
					eDist[i][j-1],
					eDist[i-1][j-1],
				)
			}
		}
	}

	// result in the entry
	return eDist[len(s1)][len(s2)]
}

func findMin(a, b, c int) int {
	// return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
	return Min(Min(a, b), c)
}

// Logic explained:
// ED[m-1][n-1] --> edit dist from s1 of len m and s2 of len n
// ED[i][j] --> edit dist between strings s1[0..i] and s2[0..j]
// 		either remove or insert ED[i-1][j], ED[i][j-1] +1
//		or alter ED[i-1][j-1] + 1 if s1[i] != s2[j]; take the min of the above
// if s1[i] == s2[j] then ED[i-1][j-1]
