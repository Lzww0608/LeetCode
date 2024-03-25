func isSubsequence(s string, t string) bool {
    i, j := 0, 0
    m, n := len(s), len(t)

    for i < m && j < n {
        if s[i] == t[j] {
            i++
        } 
        j++
    }

    if i == m {
        return true
    }
    return false
}