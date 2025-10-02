func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    
    check := func(mid int) bool {
        for i := range n {
            x := nums[i]
            if x == 0 {
                continue
            }
            f := make([]bool, x + 1)
            f[0] = true 
            for j := 0; j < mid; j++ {
                q := queries[j]
                if q[0] > i || q[1] < i {
                    continue
                }
                for y := x; y - q[2] >= 0; y-- {
                    f[y] = f[y] || f[y - q[2]]
                }
            }
            if !f[x] {
                return false
            }
        }

        return true
    }

    l, r := -1, len(queries) + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    if r == len(queries) + 1 {
        return -1
    }

    return r
}