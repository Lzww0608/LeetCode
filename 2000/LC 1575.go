const MOD int = 1_000_000_007
func countRoutes(locations []int, start int, finish int, fuel int) int {
    n := len(locations)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, fuel + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(idx, fuel int) (res int) {
        if fuel < 0 {
            return 0
        } 
        if idx == finish {
            res++
        } 
        if fuel == 0 {
            return res
        }

        p := &memo[idx][fuel]
        if *p != -1 {
            return *p
        }
        defer func(){*p = res} ()
        for j := 0; j < n; j++ {
            if idx != j {
                res += dfs(j, fuel - abs(locations[idx] - locations[j]))
            }
        }

        return res % MOD
    }

    return dfs(start, fuel)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}