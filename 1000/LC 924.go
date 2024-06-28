func minMalwareSpread(graph [][]int, initial []int) int {
    n := len(graph)
    vis := make([]bool, n)
    m := make(map[int]bool)
    for _, x := range initial {
        m[x] = true
    }
    stat := -1

    var dfs func(int) int
    dfs = func(node int) int {
        if vis[node] {
            return 0
        }
        vis[node] = true
        size := 1

        if stat != -2 && m[node] {
            if stat < 0 {
                stat = node
            } else {
                stat = -2
            }
        }

        for j := 0; j < n; j++ {
            if graph[node][j] == 1 && !vis[j] {
                size += dfs(j)
            }
        }
        return size
    }

    sort.Ints(initial)
    maxComponentSize := -1
    ans := initial[0] 


    sort.Ints(initial)

    for _, node := range initial {
        if vis[node] {
            continue
        }
        stat = -1
        componentSize := dfs(node)

        if stat >= 0 && componentSize > maxComponentSize {
            maxComponentSize = componentSize
            ans = node
        }
    }

    return ans
}
