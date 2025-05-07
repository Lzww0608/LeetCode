func maximumSum(nums []int) int64 {
    ans := 0
    n := len(nums)
    for i := 1; i <= n; i++ {
        cur := 0
        for j := 1; i * j * j <= n; j++ {
            cur += nums[i * j * j - 1]
        } 
        ans = max(ans, cur)
    }

    return int64(ans)
}