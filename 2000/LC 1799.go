func maxScore(nums []int) int {
    n := len(nums)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
        for j := i + 1; j < n; j++ {
            f[i][j] = gcd(nums[i], nums[j])
        }
    }

    dp := make([]int, 1 << n)
    for s := 1; s < (1 << n); s++ {
        if bits.OnesCount(uint(s)) & 1 == 1 {
            continue
        }
        for i := 0; i < n - 1; i++ {
            if s & (1 << i) == 0 {
                continue
            }
            for j := i + 1; j < n; j++ {
                if s & (1 << j) == 0 {
                    continue
                }
                dp[s] = max(dp[s], dp[s ^ (1 << i) ^ (1 << j)] + (bits.OnesCount(uint(s))) / 2 * f[i][j])
            }
        }
    }

    return dp[(1 << n) - 1]
}


func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}