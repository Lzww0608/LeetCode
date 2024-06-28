func lenLongestFibSubseq(arr []int) int {
    m := map[int]int{}
    n, ans := len(arr), 0
    dp := make([][]int, n)
    for i, x := range arr {
        dp[i] = make([]int, n)
        for j := i - 1; j > 0 && arr[j] * 2 > x; j-- {
            if v, ok := m[x-arr[j]]; ok && v < j {
                    dp[j][i] = max(dp[v][j] + 1, 3)
                }
                ans = max(ans, dp[j][i])
            }
        m[x] = i
    } 
    return ans
}
