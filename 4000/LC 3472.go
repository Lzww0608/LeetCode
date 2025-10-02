func longestPalindromicSubsequence(s string, k int) int {
    n := len(s)
    f := make([][][]int, n)
    for i := range f {
        f[i] = make([][]int, n)
        for j := range f[i] {
            f[i][j] = make([]int, k + 1)
        }
    }

    calc := func(c1, c2 byte) int {
        a := int(c1 - 'a')
        b := int(c2 - 'a')
        return min(abs(a - b), min(a, b) + 26 - max(a, b))
    }

    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= k; j++ {
            f[i][i][j] = 1
        }
        for j := i + 1; j < n; j++ {
            cost := calc(s[i], s[j])
            for t := 0; t <= k; t++ {
                res := max(f[i + 1][j][t], f[i][j - 1][t])
                if t >= cost {
                    res = max(res, f[i + 1][j - 1][t - cost] + 2)
                }
                f[i][j][t] = res
            }
        }
    }

    return slices.Max(f[0][n-1])
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}