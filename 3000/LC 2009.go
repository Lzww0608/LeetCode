func minOperations(a []int) int {
    sort.Ints(a)
    nums := slices.Compact(a)
    ans, l, n := 0, 0, len(a)
    for i, x := range nums {
        for nums[l] < x - n + 1 {
            l++
        }
        ans = max(ans, i - l + 1)
    }
    return n - ans
}
