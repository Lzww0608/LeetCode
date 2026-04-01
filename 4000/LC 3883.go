const N int = 5001
const MOD int = 1_000_000_007
const M int = 32
var cnt [M][]int 

func init() {
    for i := 0; i < N; i++ {
        x := 0
        for y := i; y > 0; y /= 10 {
            x += y % 10
        }
        cnt[x] = append(cnt[x], i)
    }
}

func countArrays(a []int) int {
    
    n := len(a)
    f := make([][N]int, n + 1)
    f[0][0] = 1
    for i, x := range a {
        for j := 1; j < N; j++ {
            f[i][j] = (f[i][j] + f[i][j - 1]) % MOD
        }
        if x > 31 {
            return 0
        }

        for _, y := range cnt[x] {
            f[i + 1][y] = f[i][y]
        }
    }

    ans := 0
    for _, x := range f[n] {
        ans = (ans + x) % MOD
    }

    return ans
}