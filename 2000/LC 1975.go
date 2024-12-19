
import "math"
func maxMatrixSum(matrix [][]int) int64 {
    ans, cnt := 0, 0
    m, n := len(matrix), len(matrix[0])
    mn := math.MaxInt32

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] < 0 {
                cnt++
            }
            mn = min(mn, abs(matrix[i][j]))
            ans += abs(matrix[i][j])
        }
    }

    if cnt & 1 == 0 {
        return int64(ans)
    }

    return int64(ans - 2 * mn)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}