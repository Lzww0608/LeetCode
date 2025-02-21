func maxDepthBST(order []int) int {
    n := len(order)
    fa := make([]int, n + 1)
    pos := make([]int, n + 1)
    st := []int{}

    for i, x := range order {
        pos[x] = i + 1
    }

    for x := 1; x <= n; x++ {
        i := pos[x]
        for len(st) > 0 && pos[st[len(st)-1]] > i {
            if pos[fa[st[len(st)-1]]] < i {
                fa[st[len(st)-1]] = x 
            }
            st = st[:len(st)-1]
        }

        if len(st) > 0 {
            fa[x] = st[len(st)-1]
        }
        st = append(st, x)
    }

    //fmt.Println(fa)
    for _, x := range order {
        fa[x] = 1 + fa[fa[x]]
    }

    return slices.Max(fa)
}