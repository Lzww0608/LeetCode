func canWin(currentState string) bool {
    mask := 0
    for i := range currentState {
        if currentState[i] == '+' {
            mask |= 1 << i 
        }
    }

    memo := make(map[int]bool)
    var dfs func(int) bool
    dfs = func(mask int) (ans bool) {
        if v, ok := memo[mask]; ok {
            return v
        }

        for i := 1; i < mask; i <<= 1 {
            if mask & i != 0 && mask & (i << 1) != 0 {
                ans = !dfs((mask ^ i) ^ (i << 1))
                if ans {
                    break
                }
            }
        } 
        memo[mask] = ans
        return 
    }
    
    return dfs(mask)
}