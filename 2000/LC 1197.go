func minKnightMoves(a int, b int) int {
    type pair struct {
        x, y int
    }
    
    memo := make(map[pair]int)
    var dfs func(x, y int) int
    dfs = func(x, y int) int {
        if v, ok := memo[pair{x, y}]; ok {
            return v
        }
        if x + y == 0 {
            return 0
        } else if x + y == 2 {
            return 2
        }

        res := 1 + min(dfs(abs(x - 1), abs(y - 2)), dfs(abs(x - 2), abs(y - 1)))
        memo[pair{x, y}] = res
        return res
    }

    return dfs(abs(a), abs(b))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}



