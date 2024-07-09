func incremovableSubarrayCount(nums []int) int64 {
    n, i := len(nums), 0
    for i < n - 1 && nums[i] < nums[i+1] {
        i++
    }
    if i == n - 1 {
        return int64(n * (n + 1) / 2)
    }

    ans := i + 2
    for j := n - 1; j >= 0 && (j == n - 1 || nums[j] < nums[j+1]); j-- {
        for i >= 0 && nums[i] >= nums[j] {
            i--
        }

        ans +=  i + 2
    }

    return int64(ans)
}