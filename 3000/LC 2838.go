func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
    n, m := len(heroes), len(monsters)
    id := make([]int, m)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return monsters[id[i]] < monsters[id[j]]
    })

    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    sort.Slice(p, func(i, j int) bool {
        return heroes[p[i]] < heroes[p[j]]
    })

    ans := make([]int64, n)
    sum := 0
    j := 0
    for _, i := range p {
        for j < m && heroes[i] >= monsters[id[j]] {
            sum += coins[id[j]]
            j++
        }
        ans[i] = int64(sum)
    }

    return ans
}