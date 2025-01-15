func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
    m, n := len(nums), len(changeIndices)
    if m > n {
        return -1
    }
    last := make([]int, m)

    check := func(mx int) bool {
        clear(last)
        for i, idx := range changeIndices[:mx] {
            last[idx-1] = i + 1
        }
        if slices.Contains(last, 0) {
            return false
        }

        cnt := 0
        for i, idx := range changeIndices[:mx] {
            if i + 1 == last[idx - 1] {
                if cnt -= nums[idx - 1]; cnt < 0 {
                    return false
                }
            } else {
                cnt++
            }
        }

        return true
    }


    l, r := m - 1, n + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    if r > n {
        return -1
    }
    return r
}