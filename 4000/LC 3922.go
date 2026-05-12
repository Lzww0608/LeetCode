func minFlips(s string) int {
    n := len(s)
    cnt := 0
    for i := range s {
        if s[i] == '1' {
            cnt++
        }
    }

    if cnt < 2 || cnt == n || n < 3 {
        return 0
    }

    ans := min(n - cnt, cnt - 1)
    t := 0
    if s[0] == '1' {
        t++
    }
    if s[n - 1] == '1' {
        t++
    }

    cur := cnt - t
    return min(ans, cur + 2 - t)
}