func baseNeg2(n int) string {
    if n == 0 {
        return "0"
    }
    ans := []byte{}
    for n != 0 {
        if n & 1 == 1 {
            ans = append(ans, '1')
        } else {
            ans = append(ans, '0')
        }
        n -= n & 1
        n /= -2
    }
    
    m := len(ans)
    for i, j := 0, m - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return string(ans)
}

