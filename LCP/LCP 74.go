func fieldOfGreatestBlessing(forceField [][]int) (ans int) {
    var a, b []int 
    for _, v := range forceField {
        x, y, d := v[0], v[1], v[2]
        a = append(a, 2 * x + d)
        a = append(a, 2 * x - d)
        b = append(b, 2 * y - d)
        b = append(b, 2 * y + d)
    }
    a = unique(a)
    b = unique(b)

    m, n := len(a), len(b)
    dif := make([][]int, m + 2)
    for i := range dif {
        dif[i] = make([]int, n + 2)
    }
    for _, v := range forceField {
        x, y, d := v[0], v[1], v[2]
        x1 := sort.SearchInts(a, 2 * x - d)
        x2 := sort.SearchInts(a, 2 * x + d)
        y1 := sort.SearchInts(b, 2 * y - d)
        y2 := sort.SearchInts(b, 2 * y + d)
        dif[x1 + 1][y1 + 1]++
        dif[x1 + 1][y2 + 2]--
        dif[x2 + 2][y1 + 1]--
        dif[x2 + 2][y2 + 2]++
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dif[i][j] += dif[i - 1][j] + dif[i][j - 1] - dif[i - 1][j - 1]
            ans = max(ans, dif[i][j])
        }
    }

    return 
}

func unique(a []int) []int {
    sort.Ints(a)
    n := len(a)
    i, j := 0, 0
    for i < n {
        if a[i] != a[j] {
            j++
            a[j] = a[i]
        }
        i++
    }

    return a[:j + 1]
}