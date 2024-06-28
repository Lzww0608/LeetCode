func snakesAndLadders(board [][]int) int {
    type pair struct {
        x, y int
    }
    n := len(board)

    getPos := func(x int) (int, int) {
        i := n - 1 - (x - 1) / n
        j := (x - 1) % n
        if ((x - 1) / n) & 1 == 1 {
            j = n - 1 - j
        } 
        return i, j
    }

    q := []pair{pair{1, 0}}
    vis := make([]bool, n * n + 1)
    for len(q) > 0 {
        t := q[0]
        q = q[1:]
        for k := 1; k <= 6; k++ {
            id := t.x + k
            if id > n * n {
                break
            }
            x, y := getPos(id)
            if board[x][y] > 0 {
                id = board[x][y]
            }
            if id == n * n {
                return t.y + 1
            }
            if !vis[id] {
                vis[id] = true
                q = append(q, pair{id, t.y + 1})
            }
        }
    }

    return -1
}
