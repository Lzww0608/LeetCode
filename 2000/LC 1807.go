
import "strings"
func evaluate(s string, knowledge [][]string) string {
    m := make(map[string]string)
    for _, v := range knowledge {
        m[v[0]] = v[1]
    }

    ans := strings.Builder{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            j := strings.Index(s[i:], ")") + i 
            t := s[i + 1 : j]
            if v, ok := m[t]; ok {
                ans.WriteString(v)
            } else {
                ans.WriteString("?")
            }
            i = j
        } else {
            ans.WriteByte(s[i])
        }
    }

    return ans.String()
}