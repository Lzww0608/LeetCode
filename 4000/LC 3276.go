const N int = 100
func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    nums := make([]int, N + 1)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x := grid[i][j]
            nums[x] |= 1 << i 
        }
    }

    a := [][2]int{}
    for i, x := range nums {
        if x != 0 {
            a = append(a, [2]int{i, x})
        }
        
    }

    mask := 1 << m
    f := make([][]int, len(a) + 1)
    for i := range f {
        f[i] = make([]int, mask)
    } 

    for i := 1; i < len(f); i++ {
        for s := 0; s < mask; s++ {
            f[i][s] = f[i-1][s]
            t, x := a[i-1][1], a[i-1][0]
            for ; t > 0; t &= (t - 1) {
                k := t & (-t)
                if s & k == 0 {
                    f[i][s] = max(f[i][s], f[i-1][s|k] + x)
                }
            }
        }
    }

    return f[len(a)][0]
}