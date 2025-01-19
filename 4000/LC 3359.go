func countSubmatrices(grid [][]int, k int) int64 {
    m, n := len(grid), len(grid[0])
    pre := make([]int, m)
    ans := 0
    for j := 0; j < n; j++ {
        for i := range grid {
            if k < grid[i][j] {
                pre[i] = 0
            } else if j > 0 && grid[i][j] > grid[i][j-1] {
                pre[i] = 1
            } else {
                pre[i]++
            }
        }

        st := []int{-1}
        sum := 0
        for i, x := range pre {
            for len(st) > 1 && x < pre[st[len(st)-1]] {
                h := st[len(st)-1]
                st = st[:len(st)-1]
                sum -= (h - st[len(st)-1]) * pre[h]
            }
            sum += (i - st[len(st)-1]) * x
            ans += sum
            st = append(st, i)
        }
    }

    return int64(ans)
}