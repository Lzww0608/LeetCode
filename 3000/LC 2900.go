func getLongestSubsequence(words []string, groups []int) (ans []string) {
    n := len(words)
    for i := 0; i < n; i++ {
        if i == 0 || groups[i] != groups[i-1] {
            ans = append(ans, words[i])
        }
    }

    return
}