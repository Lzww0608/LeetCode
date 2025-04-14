
import "slices"
func maxVacationDays(flights [][]int, days [][]int) int {
    n, K := len(flights), len(days[0])
    f := make([][]int, K + 1)
    for i := range f {
        f[i] = make([]int, n)
        for j := range f[i] {
            f[i][j] = -1
        }
    }

    f[0][0] = 0
    for i := 1; i <= K; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                if f[i-1][k] == -1 || flights[k][j] == 0 && j != k {
                    continue
                }
                f[i][j] = max(f[i][j], f[i-1][k] + days[j][i-1])
            }
        }
    }

    return slices.Max(f[K])
}