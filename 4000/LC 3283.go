
import (
	"math"
	"math/bits"
)
const N int = 50
func maxMoves(kx int, ky int, positions [][]int) int {
    dirs := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
    n := len(positions)
    dis := make([][N][N]int, n)

    for k, position := range positions {
        x, y := position[0], position[1]
        for i := 0; i < N; i++ {
            for j := 0; j < N; j++ {
                dis[k][i][j] = -1
            }
        }

        q := []int{x * N + y}
        dis[k][x][y] = 0
        d := 1
        for len(q) > 0 {
            for sz := len(q); sz > 0; sz-- {
                cur := q[0]
                q = q[1:]
                i, j := cur / N, cur % N
                for _, dir := range dirs {
                    ni, nj := i + dir[0], j + dir[1]
                    if ni >= 0 && ni < N && nj >= 0 && nj < N && dis[k][ni][nj] == -1 {
                        dis[k][ni][nj] = d
                        q = append(q, ni * N + nj )
                    }
                }
            }
            d++
        }
    }

    positions = append(positions, []int{kx, ky})
    m := (1 << n) - 1
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    for mask := 1; mask <= m; mask++ {
        for i, position := range positions {
            x, y := position[0], position[1]
            turn := bits.OnesCount(uint(m ^ mask)) % 2 > 0
            if turn {
                f[mask][i] = math.MaxInt
            }

            for s := uint(mask); s > 0; s &= s - 1 {
                j := bits.TrailingZeros(s)
                if turn {
                    f[mask][i] = min(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y])
                } else {
                    f[mask][i] = max(f[mask][i], f[mask ^ (1 << j)][j] + dis[j][x][y])
                }
            }
        }
    }

    return f[m][n]
}