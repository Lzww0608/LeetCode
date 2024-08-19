func findLexSmallestString(s string, a int, b int) (ans string) {
    m := map[string]bool{}
    n := len(s)

    q := []string{s}
    ans = "a"
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if cur < ans {
            ans = cur
        } 
        if m[cur] {
            continue
        }
        m[cur] = true
        tmp := []byte(cur)
        for i := 1; i < n; i += 2 {
            tmp[i] = byte('0' + (int(tmp[i] - '0') + a) % 10)
        }
        q = append(q, string(tmp))
        q = append(q, cur[n-b:] + cur[:n-b])
    }

    return
}