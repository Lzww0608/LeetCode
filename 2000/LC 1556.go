func thousandSeparator(n int) string {
    if n == 0 {
        return "0"
    }
    ans := []byte{}
    cnt := 0
    for n > 0 {
        x := n % 10 
        n /= 10
        cnt++
        ans = append(ans, byte(x + '0'))
        if cnt % 3 == 0 && n > 0 {
            ans = append(ans, '.')
        }
    }

    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return string(ans)
}