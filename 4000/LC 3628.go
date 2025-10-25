func numOfSubsequences(s string) int64 {
    ans := 0
    n := len(s)
    pre := make([][2]int, n + 1)
    suf := make([][2]int, n + 1)

    for i := range n {
        pre[i + 1] = pre[i]
        if s[i] == 'T' {
            ans += pre[i][1]
        } else if s[i] == 'C' {
            pre[i + 1][1] += pre[i][0]
        } else if s[i] == 'L' {
            pre[i + 1][0]++
        }

        j := n - i - 1
        suf[j] = suf[j + 1]
        if s[j] == 'T' {
            suf[j][1]++
        } else if s[j] == 'C' {
            suf[j][0] += suf[j + 1][1]
        }
    }
    //fmt.Println(ans)

    mx := 0
    for i := 0; i <= n; i++ {
        // L
        mx = max(mx, suf[i][0])
        // C
        mx = max(mx, pre[i][0] * suf[i][1])
        // T
        mx = max(mx, pre[i][1])
    }

    return int64(ans + mx)
}