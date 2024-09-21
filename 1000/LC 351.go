func numberOfPatterns(m int, n int) int {

    jump := [10][10]int{}
    jump[1][3], jump[3][1] = 2, 2
    jump[1][7], jump[7][1] = 4, 4
    jump[3][9], jump[9][3] = 6, 6
    jump[7][9], jump[9][7] = 8, 8
    jump[1][9], jump[9][1] = 5, 5
    jump[3][7], jump[7][3] = 5, 5
    jump[4][6], jump[6][4] = 5, 5
    jump[2][8], jump[8][2] = 5, 5
    jump[7][3], jump[3][7] = 5, 5
    jump[9][1], jump[1][9] = 5, 5

    var dfs func(current, mask, k int) int
    dfs = func(current, mask, k int) int {
        if k > n {
            return 0
        }
        var count int
        if k >= m {
            count++
        }
        for next := 1; next <= 9; next++ {
            if (mask & (1 << next)) != 0 {
                continue
            }
            if jump[current][next] != 0 && (mask & (1 << jump[current][next])) == 0 {
                continue
            }
            count += dfs(next, mask | (1 << next), k+1)
        }
        return count
    }

    var total int
    // 使用对称性减少计算量
    // 1, 3, 7, 9 对称
    total += 4 * dfs(1, 1<<1, 1)
    // 2, 4, 6, 8 对称
    total += 4 * dfs(2, 1<<2, 1)
    // 5 单独
    total += dfs(5, 1<<5, 1)

    return total
}
