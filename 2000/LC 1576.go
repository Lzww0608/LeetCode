func modifyString(s string) string {
    ans := []byte(s)
    n := len(ans)

    for i := range ans {
        if ans[i] == '?'{
            var c byte = 'a'
            for c < 'z' {
                if i > 0 && c == ans[i-1] || i + 1 < n && c == ans[i+1] {
                    c++
                } else {
                    break
                }
            }
            ans[i] = c
        }
    }

    return string(ans)
}