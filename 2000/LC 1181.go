
import (
    "sort"
    "strings"
)
func beforeAndAfterPuzzles(phrases []string) (ans []string) {
    m := make(map[string]bool)
    n := len(phrases)

    for i := 0; i < n; i++ {
        s := strings.Split(phrases[i], " ")
        for j := i + 1; j < n; j++ {
            t := strings.Split(phrases[j], " ")
            if s[len(s)-1] == t[0] {
                tmp := phrases[i]
                if len(t) > 1 {
                    tmp += " " + strings.Join(t[1:], " ")
                }
                if !m[tmp] {
                    m[tmp] = true
                    ans = append(ans, tmp)
                }
            }

            if s[0] == t[len(t)-1] {
                tmp := phrases[j]
                if len(s) > 1 {
                    tmp += " " + strings.Join(s[1:], " ")
                }
                if !m[tmp] {
                    m[tmp] = true
                    ans = append(ans, tmp)
                }
            }
        }
    }

    sort.Strings(ans)

    return
}