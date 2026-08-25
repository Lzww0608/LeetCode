func isPalindromic(s string) bool {
    t := []byte{}
    for i := range s {
        t = append(t, fmt.Sprintf("%08b", s[i])...)
    }

    n := len(t)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        if t[i] != t[j] {
            return false
        }
    }

    return true
}