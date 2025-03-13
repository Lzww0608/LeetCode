
import (
	"math"
)
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
    g := make([][]int, n)
    for i := 0; i < n; i++ {
        g[i] = make([]int, n)
        for j := 0; j < n; j++ {
            g[i][j] = -1
        }
    }

    for i, edge := range edges {
        a, b := edge[0], edge[1]
        g[a][b], g[b][a] = i, i
    }

    dijsktra := func(op, source int, f []int) (ans []int) {
        dist := make([]int, n)
        vis := make([]bool, n)
        for i := range dist {
            dist[i] = math.MaxInt32
        }
        dist[source] = 0
        for i := 0; i < n - 1; i++ {
            u := -1
            for j := 0; j < n; j++ {
                if !vis[j] && (u == -1 || dist[j] < dist[u]) {
                    u = j
                }
            }
            vis[u] = true
            for v := 0; v < n; v++ {
                if !vis[v] && g[u][v] != -1 {
                    k := g[u][v]
                    if edges[k][2] != -1 {
                        dist[v] = min(dist[v], dist[u] + edges[k][2])
                    } else if op == 0 {
                        dist[v] = min(dist[v], dist[u] + 1)
                    } else {
                        modify := target - dist[u] - f[v]
                        if modify > 0 {
                            dist[v] = min(dist[v], dist[u] + modify)
                            edges[k][2] = modify
                        } else {
                            edges[k][2] = target
                        }
                    }
                }
            }
        }

        return dist
    }

    f := dijsktra(0, destination, nil)
    if f[source] > target {
        return nil
    }
    f = dijsktra(1, source, f)
    if f[destination] != target {
        return nil
    }
    return edges
}