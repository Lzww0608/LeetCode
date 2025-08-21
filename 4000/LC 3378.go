func countComponents(nums []int, threshold int) int {
    n := len(nums)
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[ry] = rx
            n--
        }
    }

    id := make(map[int]int)
    for i, x := range nums {
        if x <= threshold {
            id[x] = i
        }
    }

    for g := 1; g <= threshold; g++ {
        mn := -1
        for x := g; x <= threshold; x += g {
            if _, ok := id[x]; ok {
                mn = x;
                break;
            }
        }

        if mn == -1 {
            continue
        }
        mx := threshold * g / mn 
        for i := mn + g; i <= mx; i += g {
            if v, ok := id[i]; ok {
                merge(id[mn], v)
            }
        }
    }

    return n
}