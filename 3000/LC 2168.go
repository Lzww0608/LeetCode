func equalDigitFrequency(s string) int {
    n := len(s)
    vis := make(map[string]bool)
    for i := 0; i < n; i++ {
        cnt := [10]int{}
        cur, sum, num := 0, 0, 0
        for j := i; j < n; j++ {
            cnt[s[j] - '0']++
            t := cnt[s[j] - '0']
            if t == 1 {
                num++
            }
            if t > sum {
                sum, cur = t, 1
            } else if t == sum {
                cur++
            }

            if cur == num {
                vis[s[i:j+1]] = true
            }
        }
    }

    return len(vis)
}