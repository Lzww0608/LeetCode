func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source)
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }

    var find func(x int) int
    find = func(x int) int {
        if x != fa[x] {
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

    for _, v := range allowedSwaps {
        x, y := v[0], v[1]
        merge(x, y)
    }

    ans := n
    m := map[int]map[int]int{}
    for i, x := range source {
        id := find(i)
        if _, ok := m[id]; !ok {
            m[id] = make(map[int]int)
        }
        m[id][x]++
    } 

    for i, x := range target {
        id := find(i)
        if _, ok := m[id]; ok && m[id][x] > 0 {
            m[id][x]--
            ans--
        }
    }

    return ans
}