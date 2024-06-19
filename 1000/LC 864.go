func shortestPathAllKeys(grid []string) int {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, { 0, -1}}
    var k, start, end int
    for i := range grid {
        for j, c := range grid[i] {
            if c >= 'a' && c <= 'z' {
                k++
            } else if c == '@' {
                start, end = i, j
            }
        }
    }

    type tuple struct {i, j, state int}
    q := []tuple{tuple{start, end, 0}}
    vis := map[tuple]bool{tuple{start, end, 0}: true}
    ans := 0
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            node := q[0]
            q = q[1:]
            i, j, state := node.i, node.j, node.state
            if state == (1 << k) - 1 {
                return ans
            }
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '#' {
                    continue
                }
                c := grid[x][y]
                if c >= 'A' && c <= 'Z' && (state >> (c - 'A')) & 1 == 0 {
                    continue
                } 
                nxt := state
                if c >= 'a' && c <= 'z' {
                    nxt |= 1 << (c - 'a')
                }
                if !vis[tuple{x, y, nxt}] {
                    vis[tuple{x, y, nxt}] = true
                    q = append(q, tuple{x, y, nxt})
                }
            }
        }
        ans++
    }

    return -1
}