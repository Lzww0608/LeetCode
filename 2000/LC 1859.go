
import "strings"
func sortSentence(str string) string {
    s := strings.Split(str, " ")
    n := len(s)
    ans := make([]string, n)

    for _, t := range s {
        m := len(t)
        x, _ := strconv.Atoi(t[m-1:])
        ans[x-1] = t[:m-1]
    }

    return strings.Join(ans, " ")
}