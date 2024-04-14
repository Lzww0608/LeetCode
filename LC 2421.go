func numberOfGoodPaths(vals []int, edges [][]int) int {
    n := len(vals)
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    id := make([]int, n)
…            }
            fa[y] = fx
        }
    }

    return ans
}