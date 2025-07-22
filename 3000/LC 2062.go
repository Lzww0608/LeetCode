
import "strings"
func countVowelSubstrings(word string) (ans int) {
    m := make(map[byte]int)

    for l, r := 0, 0; r < len(word); r++ {
        if strings.Contains("aeiou", word[r:r+1]) {
            m[word[r]] = r
            if len(m) == 5 {
                mx := r
                for _, x := range m {
                    mx = min(mx, x)
                }
                ans += mx - l + 1
            }
        } else {
            clear(m)
            l = r + 1
        }
    }

    return
}