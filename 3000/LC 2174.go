
import "math/bits"
func removeOnes(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    ans := min(m, n)

    for s := 0; s < (1 << (m * n)); s++ {
        if ans <= bits.OnesCount(uint(s)) {
            continue
        }
        tmp := make([][]int, m)
        for i := range tmp {
            tmp[i] = make([]int, n)
            copy(tmp[i], grid[i])
        }

        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if s & (1 << (i * n + j)) != 0 && tmp[i][j] == 1 {
                    for k := 0; k < n; k++ {
                        tmp[i][k] = 0
                    }
                    for k := 0; k < m; k++ {
                        tmp[k][j] = 0
                    }
                }
            }
        }

        f := true 
        for i := 0; i < m && f; i++ {
            for j := 0; j < n; j++ {
                if tmp[i][j] != 0 {
                    f = false
                    break
                }
            }
        }
        if f {
            ans = min(ans, bits.OnesCount(uint(s)))
        }
    } 

    return ans
}