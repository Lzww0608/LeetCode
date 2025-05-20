func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    m := len(queries)

    check := func(mid int) bool {
        diff := make([]int, n + 1)
        for _, v := range queries[:mid + 1] {
            diff[v[0]] += v[2]
            diff[v[1] + 1] -= v[2]
        }

        sum := 0
        for i := 0; i < n; i++ {
            sum += diff[i]
            if sum < nums[i] {
                return false
            }
        }

        return true
    }

    l, r := -1, m 
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    if r == m {
        return -1
    }

    return r + 1
}