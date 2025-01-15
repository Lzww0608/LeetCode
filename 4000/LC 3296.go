func minNumberOfSeconds(mountainHeight int, workerTimes []int) (ans int64) {
    check := func(target int) bool {
        t := mountainHeight
        for _, x := range workerTimes {
            k := target / x * 8
            t -= (int(math.Sqrt(float64(1 + k))) - 1) / 2
            if t <= 0 {
                return true
            }
        }
        return false
    }

    mx := slices.Max(workerTimes)
    h := (mountainHeight - 1) / len(workerTimes) + 1
    l, r := 0, mx * h * (h + 1) / 2 + 1

    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            ans = int64(mid)
            r = mid 
        } else {
            l = mid + 1
        }
    }

    return 
}