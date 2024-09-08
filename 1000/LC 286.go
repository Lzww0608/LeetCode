const INF int = 1 << 31 - 1
func wallsAndGates(rooms [][]int)  {
    m, n := len(rooms), len(rooms[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    q := [][]int{}


    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                q = append(q, []int{i, j, 1})
            }
        }
    }

    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        x, y, dis := cur[0], cur[1], cur[2]
        for _, dir := range dirs {
            nx, ny := x + dir[0], y + dir[1]
            if nx >= 0 && nx < m && ny >= 0 && ny < n {
                if rooms[nx][ny] > dis {
                    rooms[nx][ny] = dis
                    q = append(q, []int{nx, ny, dis + 1})
                }
            }
        }
    }


    return 
}