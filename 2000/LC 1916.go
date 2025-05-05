const MOD int = 1_000_000_007

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}
func waysToBuildRooms(prevRoom []int) int {
    n := len(prevRoom)
    g := make([][]int, n)
    a, b := 1, 1 
    for i, x := range prevRoom {
        a = a * (i + 1) % MOD;
        if x != -1 {
            g[x] = append(g[x], i)
        }
    }

    
    var dfs func(int) int 
    dfs = func(u int) (cnt int) {
        for _, v := range g[u] {
            cnt += dfs(v)
        }
        cnt++
        b = b * cnt % MOD
        return 
    }
    dfs(0)

    return a * quickPow(b, MOD - 2) % MOD
}