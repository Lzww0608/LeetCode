func addBinary(a string, b string) string {
    ans := []byte{}
    i, j := len(a) - 1, len(b) - 1
    add := 0
    for i >= 0 || j >= 0 || add > 0 {
        x := add
        if i >= 0 {
            x += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            x += int(b[j] - '0')
            j--
        }
        add = x / 2
        ans = append(ans, byte('0' + x % 2))
    }

    length := len(ans)
    for i, j := 0, length - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return string(ans)
}