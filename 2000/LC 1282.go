func groupThePeople(groupSizes []int) (ans [][]int) {
    m := make(map[int][]int)

    for i, x := range groupSizes {
        if _, ok := m[x]; !ok {
            m[x] = []int{}
        }

        m[x] = append(m[x], i)
        if len(m[x]) == x {
            ans = append(ans, m[x])
            m[x] = []int{}
        }
    }

    return 
}