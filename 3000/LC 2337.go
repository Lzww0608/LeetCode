func canChange(s string, t string) bool {
    n := len(s)
    i, j := 0, 0
    for i < n && j < n {
        for i < n && s[i] == '_' {
            i++
        }
        for j < n && t[j] == '_' {
            j++
        }
        if i < n && j < n && (s[i] != t[j] || (s[i] == 'L' && i < j) || (s[i] == 'R' && i > j)) {
            return false
        } else if i == n && j != n || j == n && i != n {
            return false
        } else {
            i++
            j++
        }
    }
    for i < n {
        if s[i] != '_' {
            return false
        }
        i++
    }
    for j < n {
        if t[j] != '_' {
            return false
        }
        j++
    }
    return true
}
