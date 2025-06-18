func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i, j int) bool {
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }
        return people[i][0] < people[j][0]
    })
    n := len(people)
    ans := make([][]int, n)
    for i := 0; i < n; i++ {
        v := people[i]
        j, cnt := 0, 0
        for cnt < v[1] {
            if ans[j] == nil || ans[j][0] >= v[0] {
                cnt++
            }
            j++
        }
        for ans[j] != nil {
            j++
        }
        ans[j] = v
    }

    return ans
}