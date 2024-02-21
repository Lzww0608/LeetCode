type pair struct {
    x, y int
}
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    ans, n := 0, len(difficulty)
    arr := []pair{}
    for i := range difficulty {
        arr = append(arr, pair{difficulty[i], profit[i]})
    }
    sort.Slice(arr, func(i, j int) bool {
        if arr[i].x != arr[j].x {
            return arr[i].x < arr[j].x
        }
        return arr[i].y > arr[j].y
    })
    sort.Ints(worker)
    i, best := 0, 0
    for _, t := range worker {
        for i < n {
            if arr[i].x > t {
                break
            }
            best = max(best, arr[i].y)
            i++
        }
        ans += best
    }
    return ans
}