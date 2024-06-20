func findChampion(n int, edges [][]int) int {
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0],e[1]
        g[b] = append(g[b], a)
    }

    ans := -1

    for i, v := range g {
        if len(v) == 0 {
            if ans == -1 {
                ans = i  
            } else {
                return -1
            }
        }
    }

    return ans
}
