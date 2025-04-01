func findCenter(edges [][]int) int {
    n := len(edges)
    in := make([]int, n + 2)
    for _, v := range edges {
        in[v[0]]++
        in[v[1]]++
    }

    for i, x := range in {
        if x == n {
            return i
        }
    }

    return -1
}