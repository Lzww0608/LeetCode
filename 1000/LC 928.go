func minMalwareSpread(graph [][]int, initial []int) int {
    n := len(graph)
    vis := make([]bool, n)
    ans, mx := -1, 0
    m := map[int]bool{}
    for _, x := range initial {
        m[x] = true
    }
    stat := 0

    var dfs func(int) int 
    dfs = func(x int) int {
        vis[x] = true
        res := 1

        for i, t := range graph[x] {
            if t == 0 {
                continue
            } 
            if m[i] {
                if stat != -2 && stat != i {
                    if stat < 0 {
                        stat = i
                    } else {
                        stat = -2
                    }
                }
            } else if !vis[i] {
                res += dfs(i)
            }
        }

        return res
    }

    cnt := make([]int, n)
    for i, x := range vis {
        if !x && !m[i] {
            stat = -1
            t := dfs(i)
            if stat >= 0 {
                cnt[stat] += t
            } 
        }
    }

    for i, x := range cnt {
        if x > mx {
            ans, mx = i, x
        }
    }
    if ans >= 0 {
        return ans
    }

    return slices.Min(initial)
}
