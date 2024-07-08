func intersectionSizeTwo(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][1] == intervals[j][1] {
            return intervals[i][0] > intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })

    a := []int{-1, -1}
    for _, v := range intervals {
        if a[len(a)-2] >= v[0] {
            continue
        } else if a[len(a)-1] < v[0] {
            a = append(a, v[1]-1)
        }
        a = append(a, v[1])
    }

    return len(a) - 2
}