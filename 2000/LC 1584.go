
import "math"
func minCostConnectPoints(points [][]int) (ans int) {
    n := len(points)

    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
    }
    dis := make([]int, n)
    for i := 0; i < n; i++ {
        dis[i] = math.MaxInt32
        for j := i + 1; j < n; j++ {
            d := abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])
            g[i][j], g[j][i] = d, d 
        }
    }

    vis := make([]bool, n)
    vis[0] = true
    for i := 1; i < n; i++ {
        dis[i] = g[0][i]
    }

    for i := 1; i < n; i++ {
        id, mn := -1, math.MaxInt32
        for j := 0; j < n; j++ {
            if !vis[j] && dis[j] < mn {
                id, mn = j, dis[j]
            }
        }

        vis[id] = true
        ans += mn 
        for j := 0; j < n; j++ {
            if !vis[j] && g[j][id] < dis[j] {
                dis[j] = g[j][id]
            }
        }
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}