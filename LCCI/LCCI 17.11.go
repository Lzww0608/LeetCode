func findClosest(words []string, word1 string, word2 string) int {
    a, b := -200000, -100000
    ans := len(words)
    for i, s := range words {
        if s == word1 {
            a = i
        } else if s == word2 {
            b = i
        }
        ans = min(ans, max(a - b, b - a))
    }

    return ans
}