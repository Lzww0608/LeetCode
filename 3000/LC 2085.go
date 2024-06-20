func countWords(words1 []string, words2 []string) int {
    m, n := map[string]int{}, map[string]int{}
    ans := 0
    for _, s := range words1 {
        m[s]++
    }
    for _, s := range words2 {
        n[s]++
    }
    for k := range m {
        if m[k] == 1 && n[k] == 1 {
            ans++
        }
    }
    return ans
}
