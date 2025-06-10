func findMaxForm(strs []string, m int, n int) (ans int) {
    sz := len(strs)
    a := make([][2]int, sz)
    for i, s := range strs {
        for j := range s {
            if s[j] == '0' {
                a[i][0]++
            } else {
                a[i][1]++
            }
        }
    }

    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    for _, v := range a {
        for i := m - v[0]; i >= 0; i-- {
            for j := n - v[1]; j >= 0; j-- {
                f[i+v[0]][j+v[1]] = max(f[i+v[0]][j+v[1]],  f[i][j] + 1)
                ans = max(ans, f[i+v[0]][j+v[1]])
            }
        }
    }

    return
}