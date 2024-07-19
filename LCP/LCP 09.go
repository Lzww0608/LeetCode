func minJump(jump []int) int {
    n := len(jump)
    vis := make([]bool, n)
    q := [][]int{{0, 1}}
    vis[0] = true
    pre := 1

    for i := 0; i < len(q); i++ {
        idx, step := q[i][0], q[i][1]
        if idx + jump[idx] >= n {
            return step
        }

        if !vis[idx + jump[idx]] {
            vis[idx + jump[idx]] = true
            q = append(q, []int{idx + jump[idx], step + 1})
        } 

        for j := pre; j < idx; j++ {
            if !vis[j] {
                vis[j] = true
                q = append(q, []int{j, step + 1})
            }
        }

        pre = max(pre, idx)
    }

    return -1
}