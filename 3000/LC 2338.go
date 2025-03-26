const MOD int = 1_000_000_007

const N int = 20_001
 
var invF [N]int
var F [N]int

var f [N][]int

var mx int = 0

func init() {
    F[0] = 1
    for i := 1; i < N; i++ {
        F[i] = F[i-1] * i % MOD
    }

    invF[N-1] = quickPow(F[N-1], MOD - 2)
    for i := N - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % MOD;
    }

    for i := 2; i < N; i++ {
        if len(f[i]) == 0 {
            for j := i; j < N; j += i {
                x, y := j, 0
                for ; x % i == 0; x /= i {
                    y++
                }
                f[j] = append(f[j], y)
                mx = max(mx, y)
            }
        }
    }
}

func idealArrays(n int, k int) (ans int) {
    for i := 1; i <= k; i++ {
        t := 1
        for _, x := range f[i] {
            t = t * comb(n + x - 1, x) % MOD
        }
        ans = (ans + t) % MOD
    }

    return
}

func comb(a, b int) int {
    if b > a {
        return 0
    }

    return F[a] * invF[b] % MOD * invF[a-b] % MOD
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