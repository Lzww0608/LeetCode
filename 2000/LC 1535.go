func getWinner(arr []int, k int) int {
    mx, t := arr[0], 0
    for i := 1; i < len(arr) && t < k; i++ {
        if arr[i] > mx {
            t = 0
            mx = arr[i]
        }
        t++
    }

    return mx
}
