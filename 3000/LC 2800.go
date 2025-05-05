
import "strings"
func minimumString(a string, b string, c string) (ans string) {
    s := []string{
        merge(merge(a, b), c),
        merge(merge(a, c), b),
        merge(merge(b, a), c),
        merge(merge(b, c), a),
        merge(merge(c, b), a),
        merge(merge(c, a), b),
    }

    for _, t := range s {
        if len(ans) == 0 || len(ans) > len(t) || len(ans) == len(t) && ans > t {
            ans = t
        }
    }

    return
}

func merge(a, b string) string {
    if strings.Contains(a, b) {
        return a
    }
    n := len(a)
    for i := 0; i < n; i++ {
        if n - i <= len(b) && a[i:] == b[:n-i] {
            return a + b[n-i:]
        }
    }

    return a + b 
}