func concatHex36(n int) string {
    ans := []byte{}
    for x := n * n * n; x > 0; x /= 36 {
        t := x % 36
        if t < 10 {
            ans = append(ans, byte('0' + t))
        } else {
            ans = append(ans, byte('A' + t - 10))
        }
    }

    for x := n * n; x > 0; x /= 16 {
        t := x % 16
        if t < 10 {
            ans = append(ans, byte('0' + t))
        } else {
            ans = append(ans, byte('A' + t - 10))
        }
    }

    slices.Reverse(ans)
    return string(ans)
}