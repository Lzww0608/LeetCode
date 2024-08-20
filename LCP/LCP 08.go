func getTriggerTime(increase [][]int, requirements [][]int) []int {
    ans := make([]int, len(requirements))
    a := make([][3]int, len(increase) + 1)
    for i, v := range increase {
        for j := 0; j < 3; j++ {
            a[i+1][j] = v[j] + a[i][j]
        }
    }

    for i, req := range requirements {
        pos := sort.Search(len(a), func(i int) bool {
            return a[i][0] >= req[0] && a[i][1] >= req[1] && a[i][2] >= req[2]
        })

        if pos == len(a) {
            ans[i] = -1
        } else {
            ans[i] = pos
        }
    }

    return ans
}