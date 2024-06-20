func smallestStringWithSwaps(str string, pairs [][]int) string {
    n := len(str)
    s := []byte(str)
    a := []int{}
    g := make([][]int, n)
    for _, v := range pairs {
        a, b := v[0], v[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    sortString := func() {
        tmp := []byte{}
        sort.Ints(a)
        for _, x := range a {
            tmp = append(tmp, s[x])
        }
        slices.Sort(tmp)
        for i, x := range a {
            s[x] = tmp[i]
        }
    }


    vis := make([]bool, n)
    var dfs func(int)
    dfs = func(i int) {
        if vis[i] {
            return 
        }
        vis[i] = true
        a = append(a, i)
        for _, x := range g[i] {
            dfs(x)
        }
    }

    for i := range s {
        if !vis[i] && len(g[i]) > 0 {
            a = nil
            dfs(i)
            sortString()
        }
    }

    return string(s)
}