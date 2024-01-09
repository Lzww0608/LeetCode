func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n, ans, l, r := len(nums), 0, 0, 0
    sum := 0
    for r < n {
        sum += nums[r]
        for l <= r && nums[r] * (r - l + 1) > sum + k {
            sum -= nums[l]
            l++
        }
        ans = max(ans, r - l + 1)
        r++
    }
    return ans
}

