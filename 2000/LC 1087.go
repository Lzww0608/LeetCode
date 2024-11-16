
import "strings"
func expand(s string) (ans []string) {
    n := len(s)

    var dfs func(int, string)
    dfs = func(idx int, cur string) {
        if idx >= n {
            ans = append(ans, cur)
            return
        }

        if s[idx] != '{' {
            dfs(idx + 1, cur + s[idx:idx+1])
        } else {
            next := strings.Index(s[idx:], "}") + idx
            t := strings.Split(s[idx+1:next], ",")
            sort.Slice(t, func(i, j int) bool {
                return t[i] < t[j]
            })
            for _, tt := range t {
                dfs(next + 1, cur + tt)
            }
        }
    }
    dfs(0, "")
    return
}