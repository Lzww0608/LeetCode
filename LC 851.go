func loudAndRich(richer [][]int, quiet []int) []int {
    n := len(quiet)
    g := make([][]int, n)
    for _, e := range richer {
        a, b := e[0], e[1]
        g[b] = append(g[b], a)
    }

    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1 // Initialize with -1 to indicate "unvisited"
    }

    var dfs func(int) int
    dfs = func(idx int) int {
        if ans[idx] != -1 {
            return ans[idx] // Return memoized result
        }
        ans[idx] = idx // Default assumption
        for _, x := range g[idx] {
            if quiet[dfs(x)] < quiet[ans[idx]] {
                ans[idx] = dfs(x) // Update with quieter person
            }
        }
        return ans[idx]
    }

    for i := 0; i < n; i++ {
        dfs(i) // Ensure all nodes are visited and memoized
    }    

    return ans
}
