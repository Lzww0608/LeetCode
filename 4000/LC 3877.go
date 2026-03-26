const N int = 20_001
func minRemovals(a []int, target int) int {
    f := make([]int, N)
    g := make([]int, N)
    for i := range f {
        f[i] = -1
    }

    f[0], g[0] = 0, 0
    for _, x := range a {
        copy(g, f)
        for i := range N / 2 {
            if g[i] == -1 {
                continue
            }
            f[i ^ x] = max(f[i ^ x], g[i] + 1)
        }
    } 

    if f[target] == -1 {
        return -1
    }
    return len(a) - f[target]
}