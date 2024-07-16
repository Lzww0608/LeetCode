func printKMoves(K int) (ans []string) {
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    pos := [4]rune{'R', 'D', 'L', 'U'}
    color := [2]rune{'_', 'X'}
    x, y := 2000, 2000
    max_x, min_x := x, x
    max_y, min_y := y, y
    cur := 0

    g := [3000][3000]int{}

    for i := 0; i < K; i++ {
        d := 1
        if g[x][y] == 1 {
            d = 3
        }
        g[x][y] ^= 1
        cur = (cur + d) % 4
        x += dirs[cur][0]
        y += dirs[cur][1]
        max_x, min_x = max(max_x, x), min(min_x, x)
        max_y, min_y = max(max_y, y), min(min_y, y)

    }

    for i := min_x; i <= max_x; i++ {
        tmp := []rune{}
        for j := min_y; j <= max_y; j++ {
            if i == x && j == y {
                tmp = append(tmp, pos[cur])
            } else {
                tmp = append(tmp, color[g[i][j]])
            }
        }
        ans = append(ans, string(tmp))
    }

    return 
}