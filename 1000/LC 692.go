
func topKFrequent(words []string, k int) (ans []string) {
    a := make(map[string]int)
    for _, word := range words {
        if _, ok := a[word]; !ok {
            ans = append(ans, word)
        }
        a[word]++
    }

    sort.Slice(ans, func(i, j int) bool {
        if a[ans[i]] == a[ans[j]] {
            return ans[i] < ans[j]
        }
        return a[ans[i]] > a[ans[j]]
    })

    return ans[:k]
}