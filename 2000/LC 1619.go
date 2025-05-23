func trimMean(arr []int) (ans float64) {
    sort.Ints(arr)
    n := len(arr)

    for i := n / 20; i < n - n / 20; i++ {
        ans += float64(arr[i])
    }

    return ans / float64(n - n / 10)
}
