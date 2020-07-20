package dp

import "math"

// LC # 174

// min health needed for knight to save princess
// grid with positive and negative health
// princess at location (m,n) with right and down movement
func calculateMinimumHP(dungeon [][]int) int {
    // init the health array
    m := len(dungeon)
    n := len(dungeon[0])

    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }

    // now iterate backward
    // from princess location dungeon(m,n)
    // dp[m-1][n-1] = 1 - dungeon[m-1][n-1]
    for j := 0; j <= n; j++ {
        dp[m][j] = math.MaxInt32
    }
    for i := 0; i <= m; i++ {
        dp[i][n] = math.MaxInt32
    }
    // good states
    dp[m][n-1] = 1
    dp[m-1][n] = 1

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            dp[i][j] = Max(1, Min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
        }
    }

    return dp[0][0]

}

// Logic: backward dp can be used to track min health needed at each location
// since we are going backward valid moves are: top, and left (opp of right, down)
// at each room/loc[i,j] min health needed is function of min health that was needed
// for loc[i+1][j] if right path is chosen, or for loc[i][j+1] if down path is chosen
// also, every loc the value has to be deducted; if the min health becomes < 0, knight dies
// so it has to be max (1, min(h[i+1][j], h[i][j+1])- dungeon[i][j])
// finally we return h[0][0]
