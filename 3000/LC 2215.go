func findDifference(a []int, b []int) [][]int {
    ans := make([][]int, 2)
    cntA := make(map[int]bool)
    cntB := make(map[int]bool)
    for _, x := range a {
        cntA[x] = true
    }
    for _, x := range b {
        if !cntA[x] && !cntB[x] {
            ans[1] = append(ans[1], x) 
        }
        cntB[x] = true
    }

    for _, x := range a {
        if !cntB[x] && cntA[x] {
            cntA[x] = false
            ans[0] = append(ans[0], x)
        }
    }

    return ans
}