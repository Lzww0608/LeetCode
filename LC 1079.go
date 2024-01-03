func numTilePossibilities(tiles string) int {
    m, alp := map[string]bool{}, map[rune]int{}
    n := len(tiles)
    for _, x := range tiles {
        alp[x]++
    }
    var dfs func(int, string) 
    dfs = func(k int, s string) {
        if k > n {
            return
        }
        m[s] = true
        for _, x := range tiles {
            if alp[x] > 0 {
                alp[x]--
                dfs(k+1, s+string(x))
                alp[x]++
            }
        }
    }
    dfs(0,"")
    return len(m)-1
}