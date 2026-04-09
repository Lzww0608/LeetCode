func maxWalls(robots []int, distance []int, walls []int) int {
    n := len(robots)
    sort.Ints(walls)
    id := make([]int, n)
    for i := range id {
        id[i] = i 
    }
    sort.Slice(id, func(i, j int) bool {
        return robots[id[i]] < robots[id[j]]
    })

    find := func(l, r int) int {
        if l > r {
            return 0
        }
        p := sort.SearchInts(walls, r + 1)
        q := sort.SearchInts(walls, l)
        return p - q
    }

    f := make([][2]int, n + 1)
    for j, i := range id {
        if j == 0 {
            f[j + 1][0] = find(max(0, robots[i] - distance[i]), robots[i] - 1)
        } else {
            t := find(max(robots[id[j - 1]] + distance[id[j - 1]] + 1, robots[i] - distance[i]), robots[i] - 1)
            tt := find(max(robots[id[j - 1]] + 1, robots[i] - distance[i]), robots[i] - 1)
            f[j + 1][0] = max(f[j][0] + tt, f[j][1] + t)

        }

        t := find(robots[i] + 1, robots[i] + distance[i])
        if j < n - 1 {
            t = find(robots[i] + 1, min(robots[i] + distance[i], robots[id[j + 1]] - 1))
        }
        f[j + 1][1] = max(f[j][0], f[j][1]) + t
    }

    ans := 0
    for _, x := range robots {
        ans += find(x, x)
    }

    return ans + max(f[n][0], f[n][1])
}