
import (
	"strings"
    "sort"
)
func addBoldTag(s string, words []string) string {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j])
    })
    n := len(s)
    vis := make([]bool, n)
    for i := 0; i < n; i++ {
        for _, t := range words {
            if strings.HasPrefix(s[i:], t) {
                for j := i; j < i + len(t); j++ {
                    vis[j] = true
                }
                break
            }
        }
    }

    builder := strings.Builder{}
    for i := 0; i < n; i++ {
        if vis[i] && (i == 0 || !vis[i-1]) {
            builder.WriteString("<b>")
        }
        builder.WriteByte(s[i])
        if vis[i] && (i == n - 1 || !vis[i+1]) {
            builder.WriteString("</b>")
        }
    }

    return builder.String()
}