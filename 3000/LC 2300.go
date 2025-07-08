func successfulPairs(spells []int, potions []int, success int64) []int {
    n, m := len(spells), len(potions)
    id := make([]int, n)
    ans := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return spells[id[i]] < spells[id[j]]
    })
    sort.Slice(potions, func(i, j int) bool {
        return potions[i] > potions[j]
    })

    j := 0
    for _, i := range id {
        for j < m && int64(potions[j] * spells[i]) >= success {
            j++
        }
        ans[i] = j
    }

    return ans
}