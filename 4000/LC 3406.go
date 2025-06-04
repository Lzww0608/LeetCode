func answerString(s string, n int) (ans string) {
    if n == 1 {
        return s
    }
    m := len(s)
    d := m - n + 1
    start := 0
    for i, j, k := 0, 1, 0; j + k < m; {
        if s[i + k] == s[j + k] {
            k++
        } else if s[i + k] < s[j + k] {
            i += k + 1 
            j = max(j, i + 1)
            k = 0
        } else {
            j += k + 1
            k = 0
        }
        start = i
    }
    ans = s[start : min(m, start + d)]

    return 
}