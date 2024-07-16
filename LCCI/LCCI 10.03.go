func search(arr []int, target int) int {
    n := len(arr)
    l, r := 0, n - 1

    for l < r {
        mid := l + ((r - l) >> 1)
        if arr[l] == arr[mid] {
            if arr[l] != target {
                l++
            } else {
                r = l
            }
        } else if arr[l] < arr[mid] {
            if arr[l] <= target && target <= arr[mid] {
                r = mid
            } else {
                l = mid + 1
            }
        } else {
            if arr[l] <= target || target <= arr[mid] {
                r = mid
            } else {
                l = mid + 1
            }
        }
    }

    if arr[l] == target {
        return l
    }
    return -1
}