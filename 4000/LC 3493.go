func numberOfComponents(properties [][]int, k int) int {
    m, n := len(properties), len(properties[0])
    fa := make([]int, m)
    for i := range fa {
        fa[i] = i
        sort.Ints(properties[i])
    }
    cnt := m 

    var find func(int) int 
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return
        }
        cnt--
        fa[rx] = ry
    }

    for i := 0; i < m; i++ {
        for j := i + 1; j < m; j++ {
            same := 0
            p, q := 0, 0 
            for p < n && q < n {
                if properties[i][p] == properties[j][q] {
                    p++
                    q++
                    same++
                } else if properties[i][p] > properties[j][q] {
                    q++
                } else {
                    p++
                }

                for p < n && p > 0 && properties[i][p] == properties[i][p-1] {
                    p++
                } 

                for q < n && q > 0 && properties[j][q] == properties[j][q-1] {
                    q++
                }
            }
            if same >= k {
                merge(i, j)
            }
        }
    }

    return cnt
}