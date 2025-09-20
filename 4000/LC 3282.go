func findMaximumScore(nums []int) int64 {
    n := len(nums)
    ans := 0

    pre := nums[0]
    for i := 0; i < n - 1; i++ {
        if nums[i] > pre {
            pre = nums[i]
        }

        ans += pre
    }

    return int64(ans)
}