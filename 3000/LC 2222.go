func numberOfWays(s string) int64 {
    ans, n := 0, len(s)
    suf := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            suf[i] = suf[i+1] + 1
        } else {
            suf[i] = suf[i+1]
        }
    }

    pre := 0
    for i, c := range s {
        if c == '0' && i - pre > 0 && suf[i+1] < n - i - 1 {
            ans += (i - pre) * (n - i - 1 - suf[i+1])
        } else if c == '1' && pre > 0 && suf[i+1] > 0 {
            ans += pre * suf[i+1]
        }

        if c == '0' {
            pre++
        }
    }
    
    return int64(ans)
}