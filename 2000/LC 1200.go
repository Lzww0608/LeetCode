func minimumAbsDifference(arr []int) (ans [][]int) {
    sort.Ints(arr)
    n := len(arr)
    mx := math.MaxInt32
    for i := 1; i < n; i++ {
        x := arr[i] - arr[i - 1]
        if x < mx {
            mx = x
            ans = [][]int{{arr[i - 1], arr[i]}}
        } else if x == mx {
            ans = append(ans, []int{arr[i - 1], arr[i]})
        }
    }

    return
}