
import "strings"
func minRemoveToMakeValid(s string) string {
    cnt := 0
    m := make(map[int]bool)
    for i := range s {
        if s[i] == '(' {
            cnt++
        } else if s[i] == ')' {
            if cnt == 0 {
                m[i] = true 
            } else {
                cnt--
            }
        }
    }
    for i := len(s) - 1; i >= 0 && cnt > 0; i-- {
        if s[i] == '(' {
            cnt--
            m[i] = true
        }
    }
    ans := strings.Builder{}
    for i := range s {
        if !m[i] {
            ans.WriteByte(s[i])
        }
    }

    return ans.String()
}