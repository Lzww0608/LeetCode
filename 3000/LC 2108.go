func firstPalindrome(words []string) string {
    check := func(k int) bool {
        n := len(words[k])
        for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
            if words[k][i] != words[k][j] {
                return false
            }
        }

        return true
    }
    
    for i := range words {
        if check(i) {
            return words[i]
        }
    }

    return ""
}