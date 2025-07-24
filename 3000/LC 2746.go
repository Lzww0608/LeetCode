const N int = 26
func minimizeConcatenatedLength(words []string) int {
    n := len(words)
    f := make([][N][N]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < N; j++ {
            for k := 0; k < N; k++ {
                f[i][j][k] = math.MaxInt32
            }
        }
    }
    x, y := int(words[0][0] - 'a'), int(words[0][len(words[0]) - 1] - 'a')
    f[0][x][y] = len(words[0])

    for i := 1; i < n; i++ {
        d := len(words[i])
        x, y = int(words[i][0] - 'a'), int(words[i][d - 1] - 'a')

        // j开头，k结尾
        for j := 0; j < N; j++ {
            for k := 0; k < N; k++ {
                // words[i] 拼在前面
                if y != j {
                    f[i][x][k] = min(f[i][x][k], f[i-1][j][k] + d)
                } else {
                    f[i][x][k] = min(f[i][x][k], f[i-1][j][k] + d - 1)
                }
                // words[i] 拼在后面
                if k != x {
                    f[i][j][y] = min(f[i][j][y], f[i-1][j][k] + d)
                } else {
                    f[i][j][y] = min(f[i][j][y], f[i-1][j][k] + d - 1)
                }
            }
        }
    }

    ans := math.MaxInt32 
    for j := 0; j < N; j++ {
        for k := 0; k < N; k++ {
            ans = min(ans, f[n-1][j][k])
        }
    }

    return ans
}