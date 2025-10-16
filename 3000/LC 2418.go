func sortPeople(names []string, heights []int) []string {
    n := len(names)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return heights[id[i]] > heights[id[j]]
    })

    ans := make([]string, n)
    for i, j := range id {
        ans[i] = names[j]
    }

    return ans
}