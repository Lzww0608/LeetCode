func smallestRangeI(nums []int, k int) int {
    mn, mx := nums[0], nums[0]
    for _, x := range nums {
        mn = min(mn, x)
        mx = max(mx, x)
    }
    
    if mx - mn <= 2 * k {
        return 0
    }
    return mx - mn - 2 * k
}
