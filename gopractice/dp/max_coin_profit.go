package dp

import "fmt"

// EPI DP Problem pp. 290-292

func maxCoinProfit(C []int) int {

	n := len(C)
	// declare profit array
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, n)
	}

	// fmt.Println(P)

	//  pre-compute the profit array diagonally
	//	dist = distnace between left index i and right index j
	for dist := 0; dist < n; dist++ {
		for i, j := 0, dist; j < n; i, j = i+1, j+1 {

			fmt.Printf("%d %d %v ", i, j, C[i:j+1])

			if i == j {
				// case only 1 element
				P[i][j] = C[i]
			} else if j == i+1 {
				// case : only 2 elements
				P[i][j] = Max(C[i], C[j])
			} else if i < n-2 && j >= 2 && i < j {
				//  case 3 or more elements
				P[i][j] = Max(
					C[i]+Min(P[i+1][j-1], P[i+2][j]),
					C[j]+Min(P[i+1][j-1], P[i][j-2]),
				)
			} else {
				// all other cases
				P[i][j] = 0
			}

			fmt.Printf("Max Profit %d\n", P[i][j])

		}
	}

	return P[0][n-1]
}

// Given a set of coinds in a row, a game is played when player 1 and player 2 can choose
// one coin from each end, objective being maximize the profit,
// question: max profit for player 1
// e.g. {5, 25, 10, 1} player 1 chooses 1, 25 player 2 chooses 10,5;
// greedy approach doesn't work, for player 1 to win he has to minimize player 2 profit

// easier apparoach: T=O(n^2) S=O(n^2)
// P(a,b) = max ( C[a] + S[a+1..b] - P[a+1,b],  C[b] + S[a...b-1] - P[a,b-1])
// where S[i..j] is the sum of coin values from index i to j both inclusive

// complex but less computation
// unravel two iterations so player 2 tries to minimize player 1's profit
// i.e. iteration 1 player tries to maximize and minimize player 2 profit
// iteration 2 player 2 tries to maximize i.e. minimize  player 1 profit
// so after 3 turns, player 1's max will be
// P(a,b) = max (
//				C[a] + min( P[a+1, b-1], P[a+2, b] )
// 				C[b] + min( P[a+1, b-1], P[a, b-2] )
// )
// where a < b, else 0
