func maxWeight(weights []int, w1 int, w2 int) int {
    f := make([][]int, w1 + 1)
    for i := range f {
        f[i] = make([]int, w2 + 1)
    }

    for _, x := range weights {
        for i := w1; i >= 0; i-- {
            for j := w2; j >= 0; j-- {
                f[i][j] = max(f[i][j], f[i][max(0, j - 1)], f[max(0, i - 1)][j])
                if i - x >= 0 {
                    f[i][j] = max(f[i][j], f[i - x][j] + x)
                }
                if j - x >= 0 {
                    f[i][j] = max(f[i][j], f[i][j - x] + x)
                }
            }
        }
    }

    return f[w1][w2]
}