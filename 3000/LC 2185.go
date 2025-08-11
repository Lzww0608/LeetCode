
import "strings"
func prefixCount(words []string, pref string) (ans int) {
    for _, word := range words {
        if strings.HasPrefix(word, pref) {
            ans++
        }
    }

    return
}