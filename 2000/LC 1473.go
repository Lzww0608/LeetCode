func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    const inf int = 1e9 + 7 

    dp := make([][]int, n+1)
    prev := make([][]int, n+1)
    for j := range dp {
        dp[j] = make([]int, target+1)
        prev[j] = make([]int, target+1)
        for k := range dp[j] {
            dp[j][k] = inf
            prev[j][k] = inf
        }
    }

    if houses[0] == 0 { 
        for j := 1; j <= n; j++ {
            dp[j][1] = cost[0][j-1] 
        }
    } else { 
        dp[houses[0]][1] = 0 
    }

    for i := 1; i < m; i++ {
        dp, prev = prev, dp
        for j := 1; j <= n; j++ {
            for k := 1; k <= target; k++ {
                dp[j][k] = inf
                if houses[i] != 0 && houses[i] != j {
                    continue 
                }

                curCost := 0
                if houses[i] == 0 { 
                    curCost = cost[i][j-1]
                }

                for p := 1; p <= n; p++ { 
                    if p != j {
                        dp[j][k] = min(dp[j][k], prev[p][k-1] + curCost) 
                    } else {
                        dp[j][k] = min(dp[j][k], prev[p][k] + curCost) 
                    }
                }
            }
        }
    }

    result := inf
    for j := 1; j <= n; j++ {
        result = min(result, dp[j][target])
    }

    if result == inf {
        return -1
    }
    return result
}
