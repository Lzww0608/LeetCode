func minArea(image [][]byte, a int, b int) int {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(image), len(image[0])
    l, r, u, d := n, 0, m, 0 

    q := [][]int{{a, b}}
    vis := make([]bool, m * n)
    vis[a * n + b] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        i, j := cur[0], cur[1]
        l = min(l, j)
        r = max(r, j)
        u = min(u, i)
        d = max(d, i)
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || image[x][y] == '0' || vis[x * n + y] {
                continue
            }
            vis[x * n + y] = true
            q = append(q, []int{x, y})
        }
        //fmt.Println(i, j, l, r, u, d)
    }
    
    return (r - l + 1) * (d - u + 1)
}