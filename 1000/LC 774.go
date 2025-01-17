const EPS float64 = 1e-6
func minmaxGasDist(stations []int, k int) float64 {
    n := len(stations)

    check := func(mid float64) bool {
        cnt := 0
        for i := 1; i < n; i++ {
            if float64(stations[i] - stations[i-1]) > mid {
                cnt += int((float64(stations[i] - stations[i-1]) - EPS) / mid)
            }
        }
        return cnt <= k
    }

    l, r := float64(0), float64(stations[n-1] - stations[0])
    for l + EPS < r {
        mid := l + ((r - l) / 2)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}