
import "math"
func assignBikes(workers [][]int, bikes [][]int) int {
    n, m := len(workers), len(bikes)

    solve := func(i, j int) int {
        return abs(workers[i][0] - bikes[j][0]) + abs(workers[i][1] - bikes[j][1])
    }

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, 1 << m)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int 
    dfs = func(i, mask int) (res int) {
        if i >= n {
            return 0
        } 
        p := &memo[i][mask]
        if *p != -1 {
            return *p
        }
        
        res = math.MaxInt32
        for j := 0; j < m; j++ {
            if (mask >> j) & 1 == 0 {
                res = min(res, solve(i, j) + dfs(i + 1, mask ^ (1 << j)))
            }
        }
        *p = res
        return
    }

    return dfs(0, 0)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}