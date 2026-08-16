func elevatorRequests(n int, start int, a []int) int64 {
    if !contain(a, start) {
        a = append(a, start)
    }
    
    m := len(a)
    sort.Ints(a)
    id := sort.SearchInts(a, start)
    f := make([][][2]int, m)
    for i := range m {
        f[i] = make([][2]int, m)
        for j := range f[i] {
            f[i][j] = [2]int{math.MaxInt / 2, math.MaxInt / 2}
        }
    }

    f[id][id] = [2]int{0, 0}
    for d := 2; d <= m; d++ {
        for l := range m - d + 1 {
            r := l + d - 1
            t := m + 1 - d 
            f[l][r][0] = min(
                f[l + 1][r][0] + (a[l + 1] - a[l]) * t,
                f[l + 1][r][1] + (a[r] - a[l]) * t)
            f[l][r][1] = min(
                f[l][r - 1][1] + (a[r] - a[r - 1]) * t,
                f[l][r - 1][0] + (a[r] - a[l]) * t)
        }
    }

    return int64(min(f[0][m - 1][0], f[0][m - 1][1]))
}

func contain(a []int, x int) bool {
    for _, y := range a {
        if x == y {
            return true
        }
    }

    return false
}