
import (
	"math"
)
const N int = 61
func isPrintable(target [][]int) bool {
    m, n := len(target), len(target[0])
    left := make([]int, N)
    right := make([]int, N)
    up := make([]int, N)
    down := make([]int, N)
    for i := 0; i < N; i++ {
        left[i] = math.MaxInt32
        right[i] = math.MinInt32
        up[i] = math.MaxInt32
        down[i] = math.MinInt32
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x := target[i][j]
            up[x] = min(up[x], i)
            down[x] = max(down[x], i)
            left[x] = min(left[x], j)
            right[x] = max(right[x], j)
        }
    }

    deg := make([]int, N)
    g := make([][]int, N)
    edges := make([][N]bool, N)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x := target[i][j]
            for color := 1; color < N; color++ {
                if left[color] <= j && right[color] >= j && up[color] <= i && down[color] >= i && color != x && !edges[color][x] {
                    deg[x]++
                    edges[color][x] = true
                    g[color] = append(g[color], x)
                } 
            }
        }
    }

    q := make([]int, 0, N)
    for i, x := range deg {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _, x := range g[cur] {
            if deg[x]--; deg[x] == 0 {
                q = append(q, x)
            }
        }
    }

    return cap(q) == 0
}