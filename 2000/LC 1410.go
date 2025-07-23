
import "strings"
func entityParser(s string) string {
    m := map[string] string {
        "&quot;": "\"",
        "&apos;": "'",  
        "&amp;": "&",
        "&gt;": ">",
        "&lt;": "<",
        "&frasl;": "/",
    }
    
    ans := strings.Builder{}
    n := len(s) 
    next:
    for i := 0; i < n; i++ {
        for d := 3; d < 8 && i + d <= n; d++ {
            t := s[i: i + d] 
            if v, ok := m[t]; ok {
                ans.WriteString(v)
                i = i + d - 1
                continue next
            }
        }
        ans.WriteByte(s[i])
    }

    return ans.String()
}