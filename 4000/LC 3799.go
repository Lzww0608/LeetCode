func wordSquares(words []string) (ans [][]string) {
    n := len(words)
    path := make([]string, 4)

    var dfs func(int) 
    dfs = func(mask int) {
        cnt := bits.OnesCount(uint(mask))
        if cnt == 4 {
            tmp := make([]string, 4)
            copy(tmp, path)
            ans = append(ans, tmp)
            return 
        }

        for i := range n {
            if (1 << i) & mask != 0 {
                continue
            }

            if cnt == 0 {
                path[0] = words[i]
                dfs(mask | (1 << i))
            } else if cnt == 1 {
                if path[0][0] == words[i][0] {
                    path[1] = words[i]
                    dfs(mask | (1 << i))
                }
            } else if cnt == 2 {
                if path[0][3] == words[i][0] {
                    path[2] = words[i]
                    dfs(mask | (1 << i))
                }
            } else {
                if path[1][3] == words[i][0] && path[2][3] == words[i][3] {
                    path[3] = words[i]
                    dfs(mask | (1 << i))
                }
            }
        }
    }
    dfs(0)

    sort.Slice(ans, func(i, j int) bool {
        for k := range 4 {
            if ans[i][k] != ans[j][k] {
                return ans[i][k] < ans[j][k]
            }
        }
        return false
    })
    return
}