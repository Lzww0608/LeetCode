func containsPattern(arr []int, m int, k int) bool {
    n := len(arr)
    cnt := 0
    for i, j := 0, m; j < n; i, j = i + 1, j + 1 {
        if arr[i] == arr[j] {
            cnt++
        } else {
            cnt = 0
        }

        if cnt >= m * (k - 1) {
            return true
        }
    }

    return false
}