func longestNiceSubarray(nums []int) (ans int) {
    n := len(nums)
    l, r := 0, 0
    or := 0
    for r < n {
        for or & nums[r] != 0 {
            or ^= nums[l]
            l++
        }
        or |= nums[r]
        ans = max(ans, r - l + 1)
        r++
    }

    return max(1, ans)
}