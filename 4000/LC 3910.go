func evenSumSubgraphs(nums []int, edges [][]int) (ans int) {
    n := len(nums)
    g := make([][]bool, n)
    for i := range g {
        g[i] = make([]bool, n)
        g[i][i] = true
    }
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u][v], g[v][u] = true, true
    }

    fa := make([]int, n)
    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry
        }
    }

    a := []int{}

    next:
    for s := 1; s < (1 << n); s++ {
        sum := 0
        a = a[:0]
        for i := range n {
            fa[i] = i
            if s & (1 << i) != 0 {
                sum += nums[i]
                a = append(a, i)
            }
        }
        if sum & 1 == 1 {
            continue
        }
        for i := range a {
            for j := i + 1; j < len(a); j++ {
                if g[a[i]][a[j]] {
                    merge(a[i], a[j])
                }
            }
        }

        pre := -1
        for i := range n {
            if s & (1 << i) != 0 {
                if pre != -1 && find(pre) != find(i) {
                    continue next
                }
                pre = i
            }
        }
        
        ans++
    }

    return
}