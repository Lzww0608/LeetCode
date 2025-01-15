func minimumTime(hens []int, grains []int) int {
    sort.Ints(hens)
    sort.Ints(grains)
    m, n := len(hens), len(grains)

    check := func(limit int) bool {
        j := 0
        for i := 0; i < m && j < n; i++ {
            x := grains[j]
            for j < n && minDis(hens[i], x, grains[j]) <= limit {
                j++
            }
        }

        return j >= n
    }

    l, r := 0, minDis(hens[0], grains[0], grains[n-1]) + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return r 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}

func minDis(x, l, r int) int {
    return abs(r - l) + min(abs(x - l), abs(x - r))
}