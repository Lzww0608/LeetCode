const MOD int = 1_000_000_007
func maxTotalValue(value []int, decay []int, m int) (ans int) {
    n := len(value)
    check := func(mid int) bool {
        t := m 
        for i := range n {
            x, y := value[i], decay[i]
            if x >= mid {
                t -= (x - mid) / y + 1 
                if t < 0 {
                    return false
                }
            }
        }
        return true
    }

    idx := 0
    if !check(0) {
        l, r := 0, slices.Max(value) + 1
        for l + 1 < r {
            mid := l + ((r - l) >> 1)
            if check(mid) {
                r = mid
            } else {
                l = mid
            }
        }

        idx = l 
    }

    for i := range n {
        x, d := value[i], decay[i]
        if x > idx {
            t := (x - idx - 1) / d + 1
            m -= t 
            ans += (x + x - d * (t - 1)) * t
        }
    }
    ans /= 2
    ans = (ans + m * idx) % MOD
    return
}