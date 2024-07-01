func statisticalResult(arr []int) []int {
    n := len(arr)
    if n == 0 {
        return nil
    }
    ans := make([]int, n + 1)
    ans[0] = 1
    t := 1
    for i, x := range arr {
        ans[i+1] = ans[i] * t
        t = x
    }

    t = arr[n-1]
    for i := n - 2; i >= 0; i-- {
        ans[i+1] *= t
        t *= arr[i]
    }

    return ans[1:]
}
