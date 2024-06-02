func lastSubstring(s string) string {
    n, i := len(s), 0
    for j, k := 1, 0; j + k < n; {
        if s[i+k] == s[j+k] {
            k++
        } else if s[i+k] < s[j+k] {
            i += k + 1
            if i >= j {
                j = i + 1
            }
            k = 0
        } else {
            j += k + 1
            k = 0
        }
    }

    return s[i:]
}
