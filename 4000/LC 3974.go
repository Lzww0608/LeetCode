func maxSum(nums []int, k int, mul int) int64 {
    sort.Ints(nums)
    n := len(nums)

    ans := 0
    for i := range k {
        ans += max(1, mul) * nums[n - i - 1]
        mul--
    }

    return int64(ans)
}