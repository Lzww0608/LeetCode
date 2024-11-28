import "slices"
func minStartingIndex(s string, pattern string) int {
    a := getZ(pattern + s)
    b := getZ(reverse(pattern) + reverse(s))
    m := len(pattern)
    slices.Reverse(b)

    for i := m; i <= len(s); i++ {
        if a[i] + b[i-1] >= m - 1 {
            return i - m
        }
    } 

    return -1
}

func getZ(s string) []int {
    n := len(s)
    z := make([]int, n)
    for i, l, r := 0, 0, 0; i < n; i++ {
        z[i] = max(min(z[i-l], r - i + 1), 0)
        for i + z[i] < n && s[z[i]] == s[i+z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }
    }

    return z
}

func reverse(s string) string {
    t := []byte(s)
    slices.Reverse(t)
    return string(t)
}