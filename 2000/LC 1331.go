func arrayRankTransform(arr []int) []int {
    n := len(arr)
    id := make([]int, n)
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return arr[id[i]] < arr[id[j]]
    })

    j, pre := 0, math.MinInt32
    for _, i := range id {
        if pre != arr[i] {
            pre = arr[i]
            j++
        } 
        ans[i] = j 
    }

    return ans
}