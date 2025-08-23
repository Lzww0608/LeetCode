const N int = 10
func minimumOperations(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    a, b, t := 0, 0, -1
    for j := 0; j < n; j++ {
        cnt := [N]int{}
        for i := 0; i < m; i++ {
            cnt[grid[i][j]]++
        }

        f1, f2, y := 0, 0, -1
        for x := 0; x < 10; x++ {
            res := a 
            if x == t {
                res = b
            }
            res += cnt[x]
            if res >= f1 {
                f2 = f1
                f1 = res 
                y = x
            } else if res > f2 {
                f2 = res
            }
        }

        a, b, t = f1, f2, y
    }

    return m * n - a
}