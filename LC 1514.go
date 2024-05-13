func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
    f := make([]float64, n)
    f[start_node] = 1.0
    for {
        k := false
        for j := 0; j < len(edges); j++ {
            if f[edges[j][0]] * succProb[j] > f[edges[j][1]] {
                f[edges[j][1]] = f[edges[j][0]] * succProb[j]
                k = true
            }
            if f[edges[j][1]] * succProb[j] > f[edges[j][0]] {
                f[edges[j][0]] = f[edges[j][1]] * succProb[j]
                k = true
            }
        }
        if !k {
            break
        }
    }

    return f[end_node]
}