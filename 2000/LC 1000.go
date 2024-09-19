
import "math"
func mergeStones(stones []int, k int) int {
    n := len(stones)
    if (n - 1) % (k - 1) != 0 {
        return -1
    }

    pre := make([]int, n + 1)
    for i, x := range stones {
        pre[i+1] = pre[i] + x
    }

    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, n)
        for j := range memo[i] {
            memo[i][j] = make([]int, k + 1)
            for p := range memo[i][j] {
                memo[i][j][p] = -1
            }
        }
    }

    var dfs func(int, int, int) int 
    dfs = func(i, j, t int) (res int) {
        p := &memo[i][j][t]
        if *p != -1 {
            return *p
        }
        defer func() {*p = res} ()

        res = math.MaxInt32
        if t == 1 {
            if i == j {
                res = 0
            } else {
                res = dfs(i, j, k) + pre[j+1] - pre[i]
            }
            return 
        }
        for m := i; m < j; m += k - 1 {
            res = min(res, dfs(i, m, 1) + dfs(m + 1, j, t - 1))
        }
        return 
    }

    return dfs(0, n - 1, 1)
}