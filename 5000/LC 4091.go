const N int = 26

func solve(s string) [][N]int {
    n := len(s)
    f := make([][N]int, n)

    for k := 1; k <= N; k++ {
        vis := make(map[byte]int)
        l := 0
        for r := 0; r < n; r++ {
            vis[s[r]]++
            for len(vis) >= k {
                if vis[s[l]]--; vis[s[l]] == 0 {
                    delete(vis, s[l])
                }
                l++
            }

            f[r][k - 1] = l 
        }
        if l == 0 {
            break
        }
    }

    for i := range f {
        for j := range N - 1 {
            if f[i][j + 1] == 0 {
                break
            } 

            f[i][j] -= f[i][j + 1]
        }
    }

    return f
}

func countMagicalPairs(s string) int64 {
    pre := solve(s)
    t := []byte(s)
    slices.Reverse(t)
    suf := solve(string(t))
    slices.Reverse(suf)

    ans := 0
    for i := range len(pre) - 1 {
        for j := range N {
            if pre[i][j] == 0 && suf[i + 1][j] == 0 {
                break
            }

            ans += pre[i][j] * suf[i + 1][j]
        }
    }

    return int64(ans)
}                                                               