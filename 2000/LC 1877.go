func minPairSum(nums []int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    for i := 0; i < n / 2; i++ {
        ans = max(ans, nums[i] + nums[n - i - 1])
    }

    return
}