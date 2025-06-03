func answerString(s string, n int) (ans string) {
    if n == 1 {
        return s
    }
    m := len(s)
    d := m - n + 1
    for i := 0; i < m; i++ {
        if ans < s[i: min(m, i + d)] {
            ans = s[i: min(m, i + d)]
        }
    }

    return
}