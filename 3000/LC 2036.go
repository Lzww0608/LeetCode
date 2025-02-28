func maximumAlternatingSubarraySum(nums []int) int64 {
    ans := nums[0]
    add, sub := nums[0], 0

    for _, x := range nums[1:] {
        add, sub = max(sub + x, x), add - x
        ans = max(ans, add, sub)
        if sub < 0 {
            sub = 0
        }
    }

    return int64(ans)
}