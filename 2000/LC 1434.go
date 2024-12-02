
import "math/bits"
const MOD int = 1_000_000_007
const N int = 40
func numberWays(hats [][]int) (ans int) {
    n := len(hats)
    m := make([]map[int]bool, N + 1)
    for i := range m {
        m[i] = make(map[int]bool)
    }

    for i, v := range hats {
        for _, x := range v {
            m[x][i] = true
        }
    }

    f := make([][]int, N + 1)
    for i := range f {
        f[i] = make([]int, 1 << n)
        for j := range f[i] {
            f[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, mask int) (res int) {
        if mask == (1 << n) - 1 {
            return 1
        } else if n - bits.OnesCount(uint(mask)) > N + 1 - i {
            return 0
        }
        if f[i][mask] != -1 {
            return f[i][mask]
        }

        res += dfs(i + 1, mask) % MOD
        for j := 0; j < n; j++ {
            if m[i][j] && (mask & (1 << j) == 0) {
                res += dfs(i + 1, mask | (1 << j))
                res %= MOD
            }
        }
        f[i][mask] = res
        return
    }

    return dfs(1, 0)
}