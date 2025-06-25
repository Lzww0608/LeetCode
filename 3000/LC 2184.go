const MOD int = 1_000_000_007
func buildWall(height int, width int, bricks []int) (ans int) {
    n := 1 << width
    masks := make([][]int, width + 1)
    masks[0] = append(masks[0], 0)
    for i := 1; i <= width; i++ {
        for _, x := range bricks {
            if x <= i {
                y := i - x 
                for _, mask := range masks[y] {
                    masks[i] = append(masks[i], mask | (1 << i))
                }
            }
        }
    }

    mask := masks[width]
    m := len(mask)
    f := make([][]int, height)
    for i := range f {
        f[i] = make([]int, m)
    }
    for j := range f[0] {
        f[0][j] = 1
    }

    for i := 1; i < height; i++ {
        for j := 0; j < m; j++ {
            for k := 0; k < m; k++ {
                if mask[j] & mask[k] == n {
                    f[i][j] = (f[i][j] + f[i-1][k]) % MOD
                }
            }
        }
    }

    for j := 0; j < m; j++ {
        ans = (ans + f[height - 1][j]) % MOD
    }

    return 
}