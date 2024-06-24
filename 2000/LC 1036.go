const N int = 1_000_000
type pair struct {
    x, y int
}
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
    n := len(blocked)
    MAX := n * (n - 1) / 2
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m := map[pair]bool{}
    for _, v := range blocked {
        m[pair{v[0], v[1]}] = true
    }

    BFS := func(start, end []int) bool {
        vis := map[pair]bool{pair{start[0], start[1]}: true}
        q := []pair{pair{start[0], start[1]}}
        for i := 0; i < len(q) && len(q) <= MAX; i++ {
            node := q[i]
            
            i, j := node.x, node.y
            for _, dir := range dirs {
                x, y := dir[0] + i, dir[1] + j
                if x < 0 || x >= N || y < 0 || y >= N || vis[pair{x, y}] || m[pair{x, y}] {
                    continue
                } 
                if x == end[0] && y == end[1] {
                    return true
                }
                vis[pair{x, y}] = true
                q = append(q, pair{x, y})
            }
        } 

        return len(q) > MAX
    }

    return BFS(source, target) && BFS(target, source)
}