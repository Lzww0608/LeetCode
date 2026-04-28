func minOperations(nums []int) int64 {
    ans := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] < 0 {
            ans += nums[i - 1] - nums[i]
        }
    }

    return int64(ans)
}