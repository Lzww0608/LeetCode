func solve(board [][]byte)  {
    m, n := len(board), len(board[0])
    dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

    var dfs func(int, int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
            return
        }
        board[i][j] = 'A'
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            dfs(x, y)
        }
    }

    for i := range board {
        dfs(i, 0)
        dfs(i, n - 1)
    }

    for j := 0; j < n; j++ {
        dfs(0, j)
        dfs(m - 1, j)
    }

    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }

    return 
}
