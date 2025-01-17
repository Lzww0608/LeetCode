func maxPower(stations []int, r int, k int) int64 {
    n := len(stations)
    pre := make([]int, n + 1)
    dif := make([]int, n)
    for i, x := range stations {
        pre[i+1] = pre[i] + x
    }

    check := func(mid int) bool {
        need := 0
        cnt := 0
        clear(dif)
        for i := 0; i < len(stations); i++ {
            cnt += dif[i]
            ll, rr := max(0, i - r), min(n - 1, i + r)
            tmp := cnt + pre[rr+1] - pre[ll]
            if tmp < mid {
                t := mid - tmp
                if need += t; need > k {
                    return false
                }
                cnt += t 
                if i + r * 2 + 1 < n {
                    dif[i + r * 2 + 1] -= t
                }
            }
        }
        return true
    }

    left, right := -1, int(2e10)
    for left + 1 < right {
        mid := left + ((right - left) >> 1)
        if check(mid) {
            left = mid
        } else {
            right = mid
        }
    }

    return int64(left)
}