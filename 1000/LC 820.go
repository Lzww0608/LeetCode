
import (
	"sort"
	"strings"
)
func minimumLengthEncoding(words []string) (ans int) {
    n := len(words)
    ans += n
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })

    for i := 0; i < n; i++ {
        ans += len(words[i])
        for j := i + 1; j < n; j++ {
            if strings.HasSuffix(words[j], words[i]) {
                ans -= len(words[i]) + 1
                break
            }
        }
    }

    return
}