var MOD int = int(1e9 + 7)
func numFactoredBinaryTrees(arr []int) int {
    m := map[int]int{}
    n := len(arr)
    dp := make([]int, n)
    sort.Ints(arr)
    for i, x := range arr {
        dp[i] = 1
        m[x] = i
        for j := i - 1; j >= 0; j-- {
            if arr[j] * arr[j] == arr[i] {
                dp[i] += dp[j] * dp[j]
            } else if arr[j] * arr[j] < arr[i] {
                break
            } else if v, ok := m[x / arr[j]]; ok && x % arr[j] == 0 {
                dp[i] += dp[j] * dp[v] * 2
            }
        }
        dp[i] %= MOD
    }
    ans := 0
    for _, x := range dp {
        ans += x
        ans %= MOD
    }
    return ans
}
