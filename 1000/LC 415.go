func addStrings(num1 string, num2 string) string {
    n, m := len(num1), len(num2)
    ans := []byte{}
    add := 0
    i, j := n - 1, m - 1
    for i >= 0 || j >= 0 || add > 0 {
        t := add
        if i >= 0 {
            t += int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            t += int(num2[j] - '0')
            j--
        }
        add = t / 10
        ans = append(ans, byte('0' + t % 10))
    }
    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return string(ans)
}
