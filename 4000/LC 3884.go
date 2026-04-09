func firstMatchingIndex(s string) int {
    n := len(s)
    for i := range (n + 1) / 2 {
        if s[i] == s[n - i - 1] {
            return i
        }
    }

    return -1
}