const MOD int = 1_000_000_007
const N int = 151
func countCoprime(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    f := make([][N]int, m + 1)
    f[0][0] = 1
    for i := range m {
        for j := range n {
            x := mat[i][j]
            for k := range N {
                y := gcd(x, k)
                f[i + 1][y] += f[i][k]
                f[i + 1][y] %= MOD
            }
        }
    }

    return f[m][1]
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}