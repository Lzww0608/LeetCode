func maximumMedianSum(nums []int) int64 {
    sort.Ints(nums)
    n := len(nums)

    ans := 0
    for i := 0; i < n / 3; i++ {
        ans += nums[n - 2 - i * 2]
    }

    return int64(ans)
}