func shortestSuperstring(w []string) string {
    n := len(w)
    mask := 1 << n
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
    }

    for i := range w {
        for j := range w {
            d := min(len(w[i]), len(w[j]))
            for k := d; k > 0; k-- {
                if w[i][len(w[i])-k:] == w[j][:k] {
                    g[i][j] = k
                    break
                }
            }
        }
    }

    f, p := make([][]int, mask), make([][]int, mask)
    for i := range f {
        f[i], p[i] = make([]int, n), make([]int, n)
    }
    for s := 0; s < mask; s++ {
        for i := range w {
            if (s >> i) & 1 == 0 {
                continue
            }
            for j := range w {
                if (s >> j) & 1 == 1 {
                    continue
                }
                if f[s | (1 << j)][j] <= f[s][i] + g[i][j] {
                    f[s | (1 << j)][j] = f[s][i] + g[i][j]
                    p[s | (1 << j)][j] = i
                }
            }
        }
    }

    mx, idx, last, status := f[mask-1][0], 0, -1, mask - 1
    for i := range w {
        if mx < f[mask-1][i] {
            mx = f[mask-1][i]
            idx = i
        }
    } 

    ans := []string{}
    for status != 0 {
        if last == -1 {
            ans = append(ans, w[idx])
        } else {
            ans = append(ans, string(w[idx][:len(w[idx]) - g[idx][last]]))
        }
        last = idx
        idx = p[status][idx]
        status ^= 1 << last
    }

    var builder strings.Builder
    for i := n - 1; i >= 0; i-- {
        builder.WriteString(ans[i])
    }
    return builder.String()
}
