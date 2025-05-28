func nextPalindrome(num string) string {
    n := len(num)
    s := nexPermutation(num[:n / 2])
    if s == nil {
        return ""
    }
    ans := make([]byte, 0, n)
    ans = append(ans, s...)
    if n & 1 == 1 {
        ans = append(ans, num[n/2])
    }
    for i := len(s) - 1; i >= 0; i-- {
        ans = append(ans, s[i])
    }

    return string(ans)
}

func nexPermutation(num string) []byte {
    n := len(num)
    s := []byte(num)
    i := n - 2
    for i >= 0 && s[i] >= s[i + 1] {
        i--
    }
    if i < 0 {
        return nil
    }
    j := n - 1
    for j >= 0 && s[i] >= s[j] {
        j--
    }
    s[i], s[j] = s[j], s[i]
    for p, q := i + 1, n - 1; p < q; p, q = p + 1, q - 1 {
        s[p], s[q] = s[q], s[p]
    }

    return s
}