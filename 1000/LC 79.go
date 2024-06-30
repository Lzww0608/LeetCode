func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    dir := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    mp := map[int]bool{}

    var dfs func(int, int, int) bool
    dfs = func(i, j, idx int) bool {
        if idx >= len(word) {
            return true
        }

        for k := 0; k < 4; k++ {
            x, y := i + dir[k][0], j + dir[k][1]
            if x < 0 || x >= m || y < 0 || y >= n || mp[x * n + y] || board[x][y] != word[idx] {
                continue
            }
            mp[x * n + y] = true
            if dfs(x, y, idx + 1) {
                return true
            }
            mp[x * n + y] = false
        }

        return false
    } 

    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] {
                mp[i * n + j] = true
                if dfs(i, j, 1) {
                    return true
                }
                mp[i * n + j] = false
            }
        }
    }
    
    return false
}
