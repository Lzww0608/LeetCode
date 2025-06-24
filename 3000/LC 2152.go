func minimumLines(points [][]int) int {
    n := len(points)
    if n <= 2 {
        return 1
    }

    m := make([][]int, n)
    for i := range m {
        m[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            mask := (1 << i) | (1 << j)
            for k := j + 1; k < n; k++ {
                if (points[i][0] - points[j][0]) * (points[i][1] - points[k][1]) == (points[i][0] - points[k][0]) * (points[i][1] - points[j][1]) {
                    mask |= 1 << k
                }
            }
            m[i][j] = mask
            m[j][i] = mask
        }
    }

    f := make([]int, 1 << n)
    for i := range f {
        f[i] = n
    }
    f[0] = 0 

    for mask := 0; mask < (1 << n); mask++ {
        if f[mask] == n {
            continue
        }

        p := 0 
        for mask & (1 << p) != 0 {
            p++
        }
        if p == n {
            break
        }

        for q := p + 1; q < n; q++ {
            if f[mask] + 1 < f[mask | m[p][q]] {
                f[mask | m[p][q]] = f[mask] + 1
            }
        }

        if f[mask] + 1 < f[mask | (1 << p)] {
            f[mask | (1 << p)] = f[mask] + 1
        }
    }

    return f[len(f)-1]
}