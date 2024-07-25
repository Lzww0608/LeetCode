func findValueOfPartition(nums []int) int {
    sort.Ints(nums)
    ans := math.MaxInt32

    for i := 1; i < len(nums); i++ {
        ans = min(ans, nums[i] - nums[i-1])
    }

    return ans
}