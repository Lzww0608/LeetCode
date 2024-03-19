func maximumScore(nums []int, k int) int {
    n := len(nums)
    ans, mn := nums[k], nums[k]

    for i, j, x := k - 1, k + 1, 0; x < n - 1; x++ {
        if j == n || (i >= 0 && nums[i] > nums[j]) {
            mn = min(mn, nums[i])
            i--
        } else {
            mn = min(mn, nums[j])
            j++
        }
        ans = max(ans, mn * (j - i - 1))
    }
    
    return ans
}