func minIncrease(nums []int) int64 {
    n := len(nums)
    suf := make([]int, n + 1)
    for i := n - 2; i >= 1; i -= 2 {
        suf[i] = suf[i + 2] + max(0, max(nums[i - 1], nums[i + 1]) - nums[i] + 1)
    }

    if n & 1 == 1 {
        return int64(suf[1])
    }

    pre := 0
    ans := suf[2]
    for i := 1; i < n - 1; i += 2 {
        pre += max(0, max(nums[i - 1], nums[i + 1]) - nums[i] + 1)
        ans = min(ans, pre + suf[i + 3])
    }

    return int64(ans)
}