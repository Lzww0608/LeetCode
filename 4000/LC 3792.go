const MOD int = 1_000_000_007
const N int = 1001

var f [N]int 

func init() {
    cur, x := 1, 2
    f[1] = 1
    for i := 2; i < N; i++ {
        pre := cur
        for range i {
            cur = cur * x % MOD
            x++
        }
        
        f[i] = f[i - 1] + cur * quickPow(pre, MOD - 2) % MOD
        f[i] %= MOD
    }
}

func sumOfBlocks(n int) int {
    return f[n]
}

func quickPow(a, r int) int {
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