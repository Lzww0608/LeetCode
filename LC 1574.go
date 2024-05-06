func findLengthOfShortestSubarray(arr []int) (ans int) {
    n := len(arr)
    r := n - 1
    for i := n - 1; i >= 0; i-- {
        if i == n - 1 || arr[i] <= arr[i+1] {
            r = i
        } else {
            break
        }
    }

    ans = r
    for i, j := 0, r; i < r; i++ {
        if i == 0 || arr[i] >= arr[i-1] {
            for j < n && arr[j] < arr[i] {
                j++
            }
            ans = min(ans, j - i - 1)
        } else {
            break
        }
    }

    return max(ans, 0)
}