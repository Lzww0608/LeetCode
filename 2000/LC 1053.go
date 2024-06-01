func prevPermOpt1(arr []int) []int {
    n := len(arr)
    for i := n - 1; i >= 1; i-- {
        if arr[i] < arr[i-1] {
            l, r := i, n - 1
            for l < r {
                mid := l + ((r - l) >> 1)
                if arr[mid] < arr[i-1] {
                    l = mid + 1
                } else {
                    r = mid
                }
            }

            for l >= n || arr[l] >= arr[i-1] || arr[l] == arr[l-1] {
                l--
            }
            arr[i-1], arr[l] = arr[l], arr[i-1]
            break
        }
    }

    return arr
}
