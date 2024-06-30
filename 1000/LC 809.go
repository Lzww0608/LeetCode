func expressiveWords(s string, words []string) int {
    ans := 0
    for _, w := range words {
        i, j := 0, 0
        for i < len(s) && j < len(w) {
            a, b := s[i], w[j]
            c, d := i, j
            if (a != b) {
                break
            }
            for i < len(s) && s[i] == a {
                i++
            }
            for j < len(w) && w[j] == b {
                j++
            }
            if (i - c < j - d || (i - c != j - d && i - c < 3)) {
                j--
                break
            }
        }
        if i == len(s) && j == len(w) {
            ans++
        }
    }
    return ans
}
