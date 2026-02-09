func reverseByType(s string) string {
    n := len(s)
    ans := make([]byte, n)
    
    a := []int{}
    b := []int{}
    for i, c := range s {
        if unicode.IsLower(c) {
            a = append(a, i)
        } else {
            b = append(b, i)
        }
    }

    for i := range a {
        ans[a[i]] = s[a[len(a) - 1 - i]]
    }
    for i := range b {
        ans[b[i]] = s[b[len(b) - 1 - i]]
    }

    return string(ans)
}