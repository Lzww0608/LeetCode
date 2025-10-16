func maxArea(coords [][]int) int64 {
    ans := 0
    l, r, u, d := math.MaxInt32, math.MinInt32, math.MinInt32, math.MaxInt32
    for _, v := range coords {
        l = min(l, v[0])
        r = max(r, v[0])
        u = max(u, v[1])
        d = min(d, v[1])
    }

    sort.Slice(coords, func(i, j int) bool {
        return coords[i][0] < coords[j][0]
    })

    n := len(coords)
    for i := 0; i < n; i++ {
        j := i
        mx, mn := math.MinInt32, math.MaxInt32
        for j < n && coords[i][0] == coords[j][0] {
            mx = max(mx, coords[j][1])
            mn = min(mn, coords[j][1])
            j++
        }

        ans = max(ans, (mx - mn) * max(coords[i][0] - l, r - coords[i][0]))
        i = j - 1
    }

    sort.Slice(coords, func(i, j int) bool {
        return coords[i][1] < coords[j][1]
    })
    for i := 0; i < n; i++ {
        j := i
        mx, mn := math.MinInt32, math.MaxInt32
        for j < n && coords[i][1] == coords[j][1] {
            mx = max(mx, coords[j][0])
            mn = min(mn, coords[j][0])
            j++
        }

        ans = max(ans, (mx - mn) * max(coords[i][1] - d, u - coords[i][1]))
        i = j - 1
    }
    
    if ans == 0 {
        return -1
    }
    return int64(ans)
}