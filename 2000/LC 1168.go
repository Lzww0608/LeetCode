func minCostToSupplyWater(n int, wells []int, pipes [][]int) (ans int) {
    fa := make([]int, n + n)
    cnt := len(fa)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) bool {
        rx, ry := find(x), find(y)
        if rx != ry {
            cnt--
            fa[rx] = ry 
            return true
        }
        return false
    }

    for i := n; i < len(fa); i++ {
        merge(i, n)
        pipes = append(pipes, []int{i + 1, i - n + 1, wells[i-n]})
    }

    sort.Slice(pipes, func(i, j int) bool {
        return pipes[i][2] < pipes[j][2]
    })

    for _, v := range pipes {
        if merge(v[0] - 1, v[1] - 1) {
            ans += v[2]
        }
        if cnt == 1 {
            return 
        }
    }
    return
}