func minInteger(num string, k int) string {
    n := len(num)
    f := make([]int, n + 1)

    update := func(x, val int) {
        for i := x; i <= n; i += i & (-i) {
            f[i] += val
        }
    }

    query := func(x int) (res int) {
        for i := x; i > 0; i -= i & (-i) {
            res += f[i]
        }

        return
    }

    pos := make([][]int, 10)
    for i, c := range num {
        x := int(c - '0')
        pos[x] = append(pos[x], i + 1)
    }

    ans := make([]byte, 0, n)
    for i := 1; i <= n; i++ {
        for j := 0; j < 10; j++ {
            if len(pos[j]) > 0 {
                l := query(n) - query(pos[j][0])
                d := pos[j][0] + l - i
                if d <= k {
                    update(pos[j][0], 1)
                    pos[j] = pos[j][1:]
                    ans = append(ans, byte('0' + j))
                    k -= d
                    break
                }
            }
        }
    }

    return string(ans)
}