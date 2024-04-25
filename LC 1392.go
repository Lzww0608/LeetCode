func longestPrefix(s string) string {
    n := len(s)
    f := make([]int, n)
    for i := range f {
        f[i] = -1
    }

    for i := 1; i < n; i++ {
        j := f[i-1]
        for j != -1 && s[j+1] != s[i] {
            j = f[j]
        }
        if s[j+1] == s[i] {
            f[i] = j + 1
        }
    }


    return s[0:f[n-1]+1]
}