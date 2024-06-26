func countQuadruplets(nums []int) (ans int64) {
    n := len(nums)
    dp := make([]int64, n)

    for i, x := range nums {
        var les int64 = 0
        for j := 0; j < i; j++ {
            if nums[j] < x {
                les++
                ans += dp[j]
            } else if nums[j] > x {
                dp[j] += les
            }
        }
    }
    return
}
