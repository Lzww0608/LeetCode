const N int = 26
func equalCountSubstrings(s string, count int) (ans int) {
    n := len(s)
    pre := make([][26]int, n + 1)
    for i := range s {
        for j := 0; j < N; j++ {
            pre[i+1][j] = pre[i][j]
        }
        x := int(s[i] - 'a')
        pre[i+1][x]++
    }

    check := func(l, r int) bool {
        for i := 0; i < N; i++ {
            x := pre[r+1][i] - pre[l][i]
            if x != 0 && x != count {
                return false
            }
        }

        return true
    }

    for i := count - 1; i < n; i++ {
        for j := i - count + 1; j >= max(0, i - N * count); j -= count {
            if check(j, i) {
                ans++
            }
        }
    }

    return
}