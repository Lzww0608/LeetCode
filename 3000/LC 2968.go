func maxFrequencyScore(nums []int, k int64) int {
    sort.Ints(nums)
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = pre[i] + x
    }

    check := func(d int) bool {
        for i := 0; i + d <= n; i++ {
            mid := i + d / 2
            l := pre[mid] - pre[i]
            r := pre[i + d] - pre[mid]
            if nums[mid] * (mid - i) - l + r - nums[mid] * (i + d - mid) <= int(k) {
                return true
            }
        }

        return false
    }

    l, r := 0, n + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}
