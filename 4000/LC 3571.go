
import "strings"
func shortestSuperstring(s1 string, s2 string) string {
    if strings.Contains(s1, s2) {
        return s1
    } else if strings.Contains(s2, s1) {
        return s2
    }
    a := s1 + s2 
    b := s2 + s1 
    for i := 0; i < len(s1); i++ {
        if s1[i:] == s2[:min(len(s2), len(s1) - i)] {
            a = s1 + s2[len(s1) - i:]
            break
        }
    } 

    for i := 0; i < len(s2); i++ {
        if s2[i:] == s1[:min(len(s1), len(s2) - i)] {
            b = s2 + s1[len(s2) - i:]
            break
        }
    }

    if len(a) < len(b) {
        return a
    }

    return b
}