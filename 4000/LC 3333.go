const MOD int = 1_000_000_007
func possibleStringCount(word string, k int) int {
    n := len(word)
    if n < k {
        return 0
    }

    cnt := []int{}
    ans, cur := 1, 0 
    for i := 0; i < n; i++ {
        cur++
        if i == n - 1 || word[i] != word[i+1] {
            if cur > 1 {
                if k > 0 {
                    cnt = append(cnt, cur - 1)
                }
                ans = ans * cur % MOD
            }
            k--
            cur = 0
        }
    }

    if k <= 0 {
        return ans
    }

    m := len(cnt)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, k)
    }
    f[0][0] = 1

    s := make([]int, k + 1)
    for i := 0; i < m; i++ {
        for j := 0; j < k; j++ {
            s[j+1] = (s[j] + f[i][j]) % MOD
        }
        for j := 0; j < k; j++ {
            f[i+1][j] = (s[j+1] - s[max(j-cnt[i], 0)])
        }
    }

    for _, x := range f[m] {
        ans -= x
    }

    return (ans % MOD + MOD) % MOD
}