const N int = 6
func minCost(n int, cost [][]int) int64 {
    ans := math.MaxInt
    
    // 0 1: 0, 0 2: 1, 1 0: 2, 1 2: 3, 2 0: 4, 2 1: 5
    // 0, 1, 2, 3, 4, 5
    f := make([][N]int, n / 2)
    f[0] = [N]int{
        cost[0][0] + cost[n - 1][1],
        cost[0][0] + cost[n - 1][2],
        cost[0][1] + cost[n - 1][0],
        cost[0][1] + cost[n - 1][2],
        cost[0][2] + cost[n - 1][0],
        cost[0][2] + cost[n - 1][1],
    }

    // 0 1: 1 2, 1 0, 2 0
    // 0 2: 1 0, 2 0, 2 1
    // 1 0: 0 1, 0 2, 2 1
    // 1 2: 0 1, 2 0, 2 1
    // 2 0: 0 1, 0 2, 1 2
    // 2 1: 0 2, 1 0, 1 2
    for i := 1; i < n / 2; i++ {
        f[i][0] = cost[i][0] + cost[n-i-1][1] + min(f[i-1][3], f[i-1][2], f[i-1][4])
        f[i][1] = cost[i][0] + cost[n-i-1][2] + min(f[i-1][2], f[i-1][4], f[i-1][5])
        f[i][2] = cost[i][1] + cost[n-i-1][0] + min(f[i-1][0], f[i-1][1], f[i-1][5])
        f[i][3] = cost[i][1] + cost[n-i-1][2] + min(f[i-1][0], f[i-1][4], f[i-1][5])
        f[i][4] = cost[i][2] + cost[n-i-1][0] + min(f[i-1][0], f[i-1][1], f[i-1][3])
        f[i][5] = cost[i][2] + cost[n-i-1][1] + min(f[i-1][1], f[i-1][2], f[i-1][3])
    }

    for i := range N {
        ans = min(ans, f[n / 2 - 1][i])
    }
    return int64(ans)
}