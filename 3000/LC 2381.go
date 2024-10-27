func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    ans := []byte(s)
    pre := make([]int, n + 1)
    for _, v := range shifts {
        if v[2] == 1 {
            pre[v[0]] += 1
            pre[v[1]+1] -= 1
        } else {
            pre[v[0]] -= 1
            pre[v[1]+1] += 1
        }
    }

    cur := 0
    for i := range ans {
        cur += pre[i]
        cur %= 26
        ans[i] = byte((int(ans[i] - 'a') + cur + 26) % 26 + 'a')
    }

    return string(ans)
}