func gardenNoAdj(n int, paths [][]int) []int {
    g := make([][]int, n + 1)
    for _, e := range paths {
        a, b := e[0] - 1, e[1] - 1
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    ans := make([]int, n)

    for i := 0; i < n; i++ {
        used := [5]bool{}
        for _, x := range g[i] {
            used[ans[x]] = true
        }
        for ans[i]++; used[ans[i]]; ans[i]++ {

        }
    }
    
    return ans
}
