const MOD int = 1_000_000_007
func sumAndMultiply(s string, queries [][]int) []int {
    n := len(s)
    sum := make([]int, n + 1)
    pre := make([]int, n + 1)
    cnt := make([]int, n + 1)
    for i := range s {
        x := int(s[i] - '0')
        sum[i + 1] = sum[i] + x
        if x != 0 {
            cnt[i + 1] = cnt[i] + 1
            pre[i + 1] = (pre[i] * 10 + x) % MOD
        } else {
            cnt[i + 1] = cnt[i]
            pre[i + 1] = pre[i]
        }
    }

    ans := make([]int, 0, len(queries))
    for _, q := range queries {
        l, r := q[0], q[1] + 1
        cur := sum[r] - sum[l]
        x := (pre[r] - pre[l] * quickPow(10, cnt[r] - cnt[l]) % MOD + MOD) % MOD * cur % MOD
        ans = append(ans, x)
    }

    return ans
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD 
        r >>= 1
    }

    return res
}