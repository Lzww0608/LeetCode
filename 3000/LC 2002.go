func maxProduct(s string) (ans int) {
    n := len(s)
    N := (1 << n) - 1
    for sub := 1; sub < N; sub++ {
        ans = max(ans, LPS(s, sub) * LPS(s, N ^ sub))
    }

    return
}

func LPS(t string, mask int) (ans int) {
    s := []byte{}
    for i := range t {
        if mask & (1 << i) != 0 {
            s = append(s, t[i])
        }
    }
    n := len(s)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
    }
    
    ans = 1
    for d := 1; d <= n; d++ {
        for i := 0; i + d <= n; i++ {
            j := i + d - 1
            if i == j {
                f[i][j] = 1
            } else if s[i] == s[j] {
                f[i][j] = f[i+1][j-1] + 2
            } else {
                f[i][j] = max(f[i+1][j], f[i][j-1])
            }
            ans = max(ans, f[i][j])
        }
    }

    return
}