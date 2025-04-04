
import "strings"
func reversePrefix(word string, ch byte) string {
    pos := strings.Index(word, string(ch))
    if pos == len(word) {
        return word
    }

    ans := make([]byte, 0, len(word))
    for i := pos; i >= 0; i-- {
        ans = append(ans, word[i])
    }
    for i := pos + 1; i < len(word); i++ {
        ans = append(ans, word[i])
    }

    return string(ans)
}