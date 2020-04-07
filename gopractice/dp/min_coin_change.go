package dp

import "fmt"

func coinChange(coins []int, amount int) int {
    //     define dp matrix for storing the optimal result
    dp := make([][]int, amount+1)
    for i := 0; i <= amount; i++ {
        dp[i] = make([]int, len(coins)+1)
    }

    //     initialize the dp array with base cases
    //     zero coins will not yield any value
    for i := range dp {
        dp[i][0] = -1
    }
    //     zero amount requires zero coins
    for j := range dp[0] {
        dp[0][j] = 0
    }

    //     memory inefficient version:
    // ietrate over  each amount
    //      for each coint value
    for a := 1; a <= amount; a++ {
        for c := 1; c <= len(coins); c++ {
            //   two cases
            var first, second int
            //  1. coin[c-1] is used
            if a-coins[c-1] >= 0 && dp[a-coins[c-1]][c] != -1 {
                first = dp[a-coins[c-1]][c] + 1
            } else {
                first = -1
            }
            //  coin[c-1] is unused
            second = dp[a][c-1]
            dp[a][c] = Min(first, second)
        }
    }
    fmt.Println(dp)
    return dp[amount][len(coins)]
}

func Min(a, b int) int {
    if a == -1 {
        return b
    }
    if b == -1 {
        return a
    }
    if a < b {
        return a
    }
    return b
}
