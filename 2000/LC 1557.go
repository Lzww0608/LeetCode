func findSmallestSetOfVertices(n int, edges [][]int) (ans []int) {
    g := make([]int, n)
    for _, e := range edges {
        g[e[1]]++
    }
    for i, x := range g {
        if x == 0 {
            ans = append(ans, i)
        }
    }

    return 
}
