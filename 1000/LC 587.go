func outerTrees(trees [][]int) [][]int {
    subtract := func(a, b []int) []int {
        return []int{a[0] - b[0], a[1] - b[1]}
    }

    cross := func(a, b []int) int {
        return a[0]*b[1] - a[1]*b[0]
    }

    getArea := func(a, b, c []int) int {
        return cross(subtract(b, a), subtract(c, a))
    }

    sort.Slice(trees, func(i, j int) bool {
        return trees[i][0] < trees[j][0] || (trees[i][0] == trees[j][0] && trees[i][1] < trees[j][1])
    })

    n := len(trees)
    if n <= 1 {
        return trees
    }

    st := make([]int, 0, n+5)
    for i := 0; i < n; i++ {
        for len(st) >= 2 && getArea(trees[st[len(st)-2]], trees[st[len(st)-1]], trees[i]) < 0 {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }

    sz := len(st)
    for i := n - 2; i >= 0; i-- {
        for len(st) > sz && getArea(trees[st[len(st)-2]], trees[st[len(st)-1]], trees[i]) < 0 {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }

    st = st[:len(st)-1]

    ans := make([][]int, 0, len(st))
    seen := map[int]bool{}
    for _, idx := range st {
        if !seen[idx] {
            ans = append(ans, trees[idx])
            seen[idx] = true
        }
    }

    return ans
}
