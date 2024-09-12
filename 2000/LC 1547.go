
import "math"
func minCost(m int, cuts []int) int {
    sort.Ints(cuts)

    n := len(cuts)
    cuts = append([]int{0}, cuts...)
    cuts = append(cuts, m)
    f := make([][]int, n + 2)
    for i := range f {
        f[i] = make([]int, n + 2)
    }

    for i := n; i >= 1; i-- {
        for j := i; j <= n; j++ {
            if i == j {
                f[i][j] = 0
            } else {
                f[i][j] = math.MaxInt32
            }

            for k := i; k <= j; k++ {
                f[i][j] = min(f[i][j], f[i][k-1] + f[k+1][j])
            }

            f[i][j] += cuts[j+1] - cuts[i-1]
        }
    }

    return f[1][n]
    
}