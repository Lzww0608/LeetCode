func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
    g := make([][]int, numCourses)
    isValid := true
    vis := make([]int, numCourses)

    for _, e := range prerequisites {
        a, b := e[0], e[1]
        g[b] = append(g[b], a)
    }

    var dfs func(int)
    dfs = func(u int) {
        vis[u] = 1
        for _, v := range g[u] {
            if vis[v] == 0 {
                dfs(v)
                if !isValid {
                    return 
                }
            } else if vis[v] == 1 {
                isValid = false
                return
            }
        }
        vis[u] = 2
        ans = append(ans, u)
    }

    for i := 0; i < numCourses; i++ {
        if vis[i] == 0 {
            dfs(i)
            if !isValid {
                return nil
            }
        }
    }

    for i, j := 0, len(ans) - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return
}