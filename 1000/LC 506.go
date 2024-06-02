func findRelativeRanks(score []int) []string {
    n := len(score)
    ans := make([]string, n)

    m := map[int]int{}
    for i, x := range score {
        m[x] = i
    }
    sort.Sort(sort.Reverse(sort.IntSlice(score)))

    for i, x := range score {
        idx := m[x]
        if i == 0 {
            ans[idx] = "Gold Medal" 
        } else if i == 1 {
            ans[idx] = "Silver Medal"
        } else if i == 2 {
            ans[idx] = "Bronze Medal"
        } else {
            ans[idx] = strconv.FormatInt(int64(i + 1), 10)
        }
    }

    return ans
}
