func canMakeSubsequence(s string, t string) bool {
    m, n := len(s), len(t)
    if n > m {
        return false
    }

    i, j := 0, 0
    for i < m && j < n {
        if s[i] == t[j] || byte('a' + int(s[i] - 'a' + 1) % 26) == t[j] {
            j++
        } 
        i++
    }

    return j >= n
}