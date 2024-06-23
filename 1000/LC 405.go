func toHex(n int) string {
    ans := []byte{}
    num := uint32(n)
    for num != 0 {
        t := num & 0xf
        num >>= 4
        if t < 10 {
            ans = append(ans, byte('0' + t))
        } else {
            ans = append(ans, byte('a' + t - 10))
        }
    }
    if len(ans) == 0 {
        return "0"
    }

    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    } 
    return string(ans)
}
