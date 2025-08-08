func minFlipsMonoIncr(s string) int {
    n := len(s)
    suf := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            suf[i] = suf[i + 1] + 1
        } else {
            suf[i] = suf[i + 1]
        }
    }

    pre := 0
    ans := suf[0]
    for i := range s {
        if s[i] == '1' {
            pre++
        }
        ans = min(ans, pre + suf[i + 1])
    }

    return ans
}