
import "strings"
func countPrefixes(words []string, s string) (ans int) {
    for _, t := range words {
        if strings.HasPrefix(s, t) {
            ans++
        }
    }

    return 
}