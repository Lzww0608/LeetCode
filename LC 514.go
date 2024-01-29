func findRotateSteps(s string, t string) int {
    n, m := len(s), len(t)

    pos := make([][]int, 26)
    for i, c := range s {
        pos[c-'a'] = append(pos[c-'a'], i)
    }

    mem := make([][]int, n)
    for i := range mem {
        mem[i] = make([]int, m)
        for j := range mem[i] {
            mem[i][j] = -1
        }
    }

    var dfs func(x, y int) int
    dfs = func(x, y int) int {
        if y == m {
            return 0
        }
        if mem[x][y] != -1 {
            return mem[x][y]
        }
        cur := t[y]
        num := cur - 'a'
        res := int(0x7f7f7f7f)
        for _, target := range pos[num] {
            d1 := abs(x - target)
            d2 := n - d1
            res = min(res, min(d1, d2) + dfs(target, y + 1))
        }
        mem[x][y] = res
        return res
    }
    return m + dfs(0, 0)
}


func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}