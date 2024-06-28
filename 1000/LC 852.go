func peakIndexInMountainArray(arr []int) int {
    n := len(arr)
    l, r := 0, n - 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
            return mid
        } else if arr[mid] > arr[mid-1] {
            l = mid
        } else {
            r = mid
        }
    }
    return l
}
