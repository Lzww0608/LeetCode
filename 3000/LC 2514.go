const MOD int = 1_000_000_007

const N int = 100_001
var factor [N]int 

func init() {
    factor[0] = 1
    for i := 1; i < N; i++ {
        factor[i] = factor[i-1] * i % MOD
    }
}

func quickPow(a, r int) int {
    if a == 1 {
        return 1
    }
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        r >>= 1
        a = a * a % MOD
    }

    return res 
}

func modInverse(a, b int) int {
    return a * quickPow(b, MOD - 2) % MOD
}

func countAnagrams(s string) int {
    t := strings.Split(s, " ")
    ans := 1

    for _, v := range t {
        n := len(v)
        y := 1
        cnt := make(map[byte]int)
        for i := range v {
            cnt[v[i]]++
        }
        for _, x := range cnt {
            y = y * factor[x] % MOD
        }
        ans = ans * modInverse(factor[n], y) % MOD
    }

    return ans
}