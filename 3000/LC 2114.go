import "strings"
func mostWordsFound(sentences []string) (ans int) {
    for _, sentence := range sentences {
        ans = max(ans, strings.Count(sentence, " ") + 1)
    }

    return 
}