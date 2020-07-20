package strs

// LC #44

// e.g. text: "abcabd" p: "a*d" --> true
// 		p := "a*b?" --> true, p := "a*c?d" --> false

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	// define m x n dp array for len of text, pattern
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	// base case: nil string nil pattern is  match
	dp[0][0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// start matching from 1st char
			if i > 0 && j > 0 {
				if s[i-1] == p[j-1] || p[j-1] == '?' {
					dp[i][j] = dp[i-1][j-1]
					continue
				}
			}
			// case when pattern is star
			if j > 0 && p[j-1] == '*' {
				// string may be nil
				if i > 0 {
					dp[i][j] = dp[i][j-1] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[n][m]

}

// Logic:
// Two ways of doing this, suffix based and prefix based expansion
// here's the prefix way of expanding
// In the pattern '*' means 0 or more chars '?' refers to 1 char
// Thus we have 3 cases:
// 	i. chars match between text and pattern, t[i] == p[i] or
//  	p[i] == '?', in this case we skip ahead by 1 char
//  ii. pattern has a '*', this case we have 2 options:
// 		a. '*' corresponds to no chars, i.e. we match with test by skipping the '*' in pattern,
// 		b. '* ' corresponds to 1 or more chars in text, in this case we match 1 char in the text
//  	and recusively check if text[i-1] is a match or not? similarly, text[i-2] is automatically checked
// finally, we have the result in dp[n][m] i.e. for the entire len of text and pattern
