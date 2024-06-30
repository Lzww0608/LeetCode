func addBinary(a string, b string) string {
    if len(a) < len(b) {
        return addBinary(b, a)
    }
    n, m := len(a), len(b)
    i, j := n - 1, m - 1
    
    ans := make([]byte, n)
    add := 0
    for i >= 0 && j >= 0 {
        x, y := int(a[i] - '0'), int(b[j] - '0')
        t := (x + y + add) % 2
        add = (x + y + add) / 2
        ans[i] = byte('0' + t)
        i--
        j--
    }

    for i >= 0 {
        x := int(a[i] - '0')
        t := (x + add) % 2
        add = (x + add) / 2
        ans[i] = byte('0' + t)
        i--
    }

    if add == 1 {
        ans = append([]byte{'1'}, ans...)
    }
    return string(ans)
}
