const EPS float64 = 1e-5
func equalizeWater(buckets []int, loss int) float64 {
    mx := slices.Max(buckets)
    sort.Slice(buckets, func(i, j int) bool {
        return buckets[i] > buckets[j]
    })

    check := func(limit float64) bool {
        var a, b float64
        for _, x := range buckets {
            y := float64(x)
            if y >= limit {
                a += y - limit
            } else {
                b += (limit - y) * 100.0 / float64(100 - loss)
            }
        }
        return a >= b
    }

    l, r := float64(0), float64(mx + 1) 
    for l + EPS < r {
        mid := l + ((r - l) / 2)
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}