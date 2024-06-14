const P int = 1313131
func longestDupSubstring(s string) (ans string) {
    n := len(s)
    h := make([]int, n + 5)
    p := make([]int, n + 5)
    p[0] = 1
    for i := range s {
        p[i+1] = p[i] * P
        h[i+1] = h[i] * P + int(s[i] - 'a')
    }

    check := func(k int) string {
        m := map[int]bool{}
        for i := 1; i + k - 1 <= n; i++ {
            j := i + k - 1
            cur := h[j] - h[i-1] * p[j-i+1]
            if m[cur] {
                return s[i-1:j]
            }
            m[cur] = true
        }
        return ""
    }

    l, r := 0, n
    for l < r {
        mid := l + ((r - l) >> 1)
        t := check(mid)
        if len(t) != 0 {
            l = mid + 1
            ans = t
        } else {
            r = mid
        }
    }

    return
}