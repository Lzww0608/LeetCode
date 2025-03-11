const MOD int = 1_000_000_007
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
    n := len(jump)
    memo := make(map[int]int)

    var dfs func(int) int
    dfs = func(cur int) (res int) {
        if cur <= 1 {
            return cur * inc
        }

        if v, ok := memo[cur]; ok {
            return v 
        }
        
        res = cur * inc 
        for i := 0; i < n; i++ {
            res = min(res, dfs(cur / jump[i]) + cost[i] + cur % jump[i] * inc)
            if cur % jump[i] != 0 {
                res = min(res, dfs(cur / jump[i] + 1) + cost[i] + dec * (jump[i] - cur % jump[i]))
            }
        }
        memo[cur] = res 
        return res
    }

    return dfs(target) % MOD
}