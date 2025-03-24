const MOD int = 1_000_000_007
const N int = 10_105

var F [N]int
var invF [N]int
var LPF [N]int

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
        if LPF[i] == 0 {
            for j := i; j < N; j += i {
                if LPF[j] == 0 {
                    LPF[j] = i
                }
            }
        }
    }
}

func comb(a, b int) int {
    return F[a] * invF[b] % MOD * invF[a-b] % MOD
}


func waysToFillArray(queries [][]int) []int {
    m := len(queries)
    ans := make([]int, m)

    for i, v := range queries {
        n, k := v[0], v[1]
        cur := 1
        for k > 1 {
            cnt, t := 1, LPF[k]
            for k /= t; k % t == 0; k /= t {
                cnt++
            }
            cur = (cur * comb(n + cnt - 1, cnt)) % MOD
        }
        ans[i] = cur
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