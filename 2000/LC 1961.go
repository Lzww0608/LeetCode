func isPrefixString(s string, words []string) bool {
    n := len(s)
    i := 0
    for _, word := range words {
        for j := range word {
            if i >= n {
                return false
            }
            if s[i] != word[j] {
                return false
            }
            i++
        }
        if i >= n {
            return true
        }
    }

    return false
}