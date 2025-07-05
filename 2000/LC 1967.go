import "strings"
func numOfStrings(patterns []string, word string) (ans int) {
    for _, s := range patterns {
        if strings.Contains(word, s) {
            ans++
        }
    }

    return
}