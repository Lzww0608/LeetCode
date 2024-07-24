func peakIndexInMountainArray(arr []int) int {
    l, r := 0, len(arr)
    for l < r {
        mid := l + ((r - l) >> 1)
        if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
            return mid
        } else if arr[mid] > arr[mid-1] {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return l
}