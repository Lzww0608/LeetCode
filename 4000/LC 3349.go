func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    f := make([]int, n)
    f[0] = 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            f[i] = f[i - 1] + 1
        } else {
            f[i] = 1
        }
    }

    for i := 0; i + k * 2 <= n; i++ {
        if f[i + k - 1] - f[i] == k - 1 && f[i + k * 2 - 1] - f[i + k] == k - 1 {
            return true
        }
    }

    return false
}