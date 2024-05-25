func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    fa := make([]int, n + 1)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    for _, e := range edges {
        a, b := find(e[0]), find(e[1])
        if a != b {
            fa[a] = b 
        } else {
            return []int{e[0], e[1]}
        }
    }

    return nil
}