
import "strconv"
func compressString(s string) string {
    ans := []byte{}
    cnt := 1
    n := len(s)
    for i := 1; i <= n; i++ {
        if i < n && s[i] == s[i-1] {
            cnt++
        } else {
            t := strconv.Itoa(cnt)
            ans = append(ans, s[i-1])
            ans = append(ans, t...)
            cnt = 1
        }
    }
    if len(ans) < n {
        return string(ans)
    }

    return s
}