import "strings"
func indexPairs(text string, words []string) (ans [][]int) {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    for i := range text {
        for _, s := range words {
            if i + len(s) <= len(text) && text[i:i+len(s)] == s {
                ans = append(ans, []int{i,i + len(s) - 1})
            }
        }
    }

    return
}